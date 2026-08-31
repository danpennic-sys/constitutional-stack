package store

import (
	"yourmodule/compaction"
	"yourmodule/envelope"
)

// SnapshotStore is the minimal constitutional interface.
// No inference. No mutation of envelopes.
type SnapshotStore interface {
	SaveSnapshot(*compaction.Snapshot) (string, error)
	GetSnapshot(id string) (*compaction.Snapshot, error)
	LoadSegment(startID, endID string) ([]envelope.Envelope, error)
}
