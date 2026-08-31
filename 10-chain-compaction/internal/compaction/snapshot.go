package compaction

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"yourmodule/crypto"
	"yourmodule/envelope"
)

type Snapshot struct {
	Type                 string `json:"type"`
	SegmentStart         string `json:"segment_start"`
	SegmentEnd           string `json:"segment_end"`
	RootHash             string `json:"root_hash"`
	CompactionDigest     string `json:"compaction_digest"`
	VerifierID           string `json:"verifier_id"`
	PublicKeyFingerprint string `json:"public_key_fingerprint"`
	Version              string `json:"version"`
	Timestamp            string `json:"timestamp"`
	Signature            string `json:"signature"`
}

func ComputeRootHash(envelopes []envelope.Envelope) (string, error) {
	h := sha256.New()
	for _, env := range envelopes {
		canonical, err := env.CanonicalBytes()
		if err != nil {
			return "", fmt.Errorf("canonical bytes for %s: %w", env.ID, err)
		}
		h.Write(canonical)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func ComputeCompactionDigest(rootHash, start, end, ts string) string {
	h := sha256.New()
	h.Write([]byte(rootHash))
	h.Write([]byte(start))
	h.Write([]byte(end))
	h.Write([]byte(ts))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateSnapshot(
	segment []envelope.Envelope,
	verifierID, version, fingerprint string,
	signer crypto.Signer,
) (*Snapshot, error) {
	if len(segment) == 0 {
		return nil, fmt.Errorf("empty segment")
	}

	start := segment[0].ID
	end := segment[len(segment)-1].ID

	rootHash, err := ComputeRootHash(segment)
	if err != nil {
		return nil, err
	}

	ts := time.Now().UTC().Format(time.RFC3339)
	digest := ComputeCompactionDigest(rootHash, start, end, ts)

	snap := &Snapshot{
		Type:                 "compaction_snapshot",
		SegmentStart:         start,
		SegmentEnd:           end,
		RootHash:             rootHash,
		CompactionDigest:     digest,
		VerifierID:           verifierID,
		PublicKeyFingerprint: fingerprint,
		Version:              version,
		Timestamp:            ts,
	}

	payload, err := json.Marshal(struct {
		Type                 string `json:"type"`
		SegmentStart         string `json:"segment_start"`
		SegmentEnd           string `json:"segment_end"`
		RootHash             string `json:"root_hash"`
		CompactionDigest     string `json:"compaction_digest"`
		VerifierID           string `json:"verifier_id"`
		PublicKeyFingerprint string `json:"public_key_fingerprint"`
		Version              string `json:"version"`
		Timestamp            string `json:"timestamp"`
	}{
		Type:                 snap.Type,
		SegmentStart:         snap.SegmentStart,
		SegmentEnd:           snap.SegmentEnd,
		RootHash:             snap.RootHash,
		CompactionDigest:     snap.CompactionDigest,
		VerifierID:           snap.VerifierID,
		PublicKeyFingerprint: snap.PublicKeyFingerprint,
		Version:              snap.Version,
		Timestamp:            snap.Timestamp,
	})
	if err != nil {
		return nil, err
	}

	sig, err := signer.Sign(payload)
	if err != nil {
		return nil, err
	}
	snap.Signature = base64.StdEncoding.EncodeToString(sig)

	return snap, nil
}

func VerifySnapshot(snap *Snapshot, verifier crypto.Verifier) error {
	if snap.Type != "compaction_snapshot" {
		return fmt.Errorf("unexpected type %q", snap.Type)
	}

	payload, err := json.Marshal(struct {
		Type                 string `json:"type"`
		SegmentStart         string `json:"segment_start"`
		SegmentEnd           string `json:"segment_end"`
		RootHash             string `json:"root_hash"`
		CompactionDigest     string `json:"compaction_digest"`
		VerifierID           string `json:"verifier_id"`
		PublicKeyFingerprint string `json:"public_key_fingerprint"`
		Version              string `json:"version"`
		Timestamp            string `json:"timestamp"`
	}{
		Type:                 snap.Type,
		SegmentStart:         snap.SegmentStart,
		SegmentEnd:           snap.SegmentEnd,
		RootHash:             snap.RootHash,
		CompactionDigest:     snap.CompactionDigest,
		VerifierID:           snap.VerifierID,
		PublicKeyFingerprint: snap.PublicKeyFingerprint,
		Version:              snap.Version,
		Timestamp:            snap.Timestamp,
	})
	if err != nil {
		return err
	}

	sig, err := base64.StdEncoding.DecodeString(snap.Signature)
	if err != nil {
		return fmt.Errorf("signature decode: %w", err)
	}

	if err := verifier.Verify(payload, sig, snap.PublicKeyFingerprint); err != nil {
		return fmt.Errorf("signature: %w", err)
	}

	expectedDigest := ComputeCompactionDigest(snap.RootHash, snap.SegmentStart, snap.SegmentEnd, snap.Timestamp)
	if expectedDigest != snap.CompactionDigest {
		return fmt.Errorf("compaction_digest mismatch")
	}
	return nil
}
