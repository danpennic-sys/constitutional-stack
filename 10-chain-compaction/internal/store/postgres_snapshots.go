package store

import (
	"database/sql"
	"encoding/base64"
	"fmt"

	"yourmodule/compaction"
	"yourmodule/envelope"
)

type PostgresSnapshotStore struct {
	db *sql.DB
}

func NewPostgresSnapshotStore(db *sql.DB) *PostgresSnapshotStore {
	return &PostgresSnapshotStore{db: db}
}

func (s *PostgresSnapshotStore) SaveSnapshot(snap *compaction.Snapshot) (string, error) {
	var id string
	sig, err := base64.StdEncoding.DecodeString(snap.Signature)
	if err != nil {
		return "", fmt.Errorf("decode signature: %w", err)
	}

	err = s.db.QueryRow(`
		INSERT INTO snapshots (
			type,
			segment_start,
			segment_end,
			root_hash,
			compaction_digest,
			verifier_id,
			version,
			timestamp,
			signature
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING snapshot_id
	`,
		snap.Type,
		snap.SegmentStart,
		snap.SegmentEnd,
		snap.RootHash,
		snap.CompactionDigest,
		snap.VerifierID,
		snap.Version,
		snap.Timestamp,
		sig,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("save snapshot: %w", err)
	}
	return id, nil
}

func (s *PostgresSnapshotStore) GetSnapshot(id string) (*compaction.Snapshot, error) {
	row := s.db.QueryRow(`
		SELECT
			type,
			segment_start,
			segment_end,
			root_hash,
			compaction_digest,
			verifier_id,
			version,
			timestamp,
			signature
		FROM snapshots
		WHERE snapshot_id = $1
	`, id)

	var snap compaction.Snapshot
	var sig []byte

	if err := row.Scan(
		&snap.Type,
		&snap.SegmentStart,
		&snap.SegmentEnd,
		&snap.RootHash,
		&snap.CompactionDigest,
		&snap.VerifierID,
		&snap.Version,
		&snap.Timestamp,
		&sig,
	); err != nil {
		return nil, fmt.Errorf("get snapshot: %w", err)
	}

	snap.Signature = base64.StdEncoding.EncodeToString(sig)
	return &snap, nil
}

func (s *PostgresSnapshotStore) LoadSegment(startID, endID string) ([]envelope.Envelope, error) {
	rows, err := s.db.Query(`
		SELECT event_id, schema_version, ts, type, payload, payload_digest, lineage
		FROM envelopes
		WHERE ts >= (SELECT ts FROM envelopes WHERE event_id = $1)
		  AND ts <= (SELECT ts FROM envelopes WHERE event_id = $2)
		ORDER BY ts ASC
	`, startID, endID)
	if err != nil {
		return nil, fmt.Errorf("load segment: %w", err)
	}
	defer rows.Close()

	var out []envelope.Envelope
	for rows.Next() {
		var env envelope.Envelope
		if err := rows.Scan(
			&env.ID,
			&env.SchemaVersion,
			&env.Ts,
			&env.Type,
			&env.Payload,
			&env.PayloadDigest,
			&env.Lineage,
		); err != nil {
			return nil, err
		}
		out = append(out, env)
	}
	return out, nil
}
