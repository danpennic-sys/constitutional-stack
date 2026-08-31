package archive

import (
	"fmt"

	"yourmodule/store"
)

// Mover relocates a compacted segment to the archive tier.
// Constitutional rule: never mutate or delete envelopes or proofs — only relocate.
type Mover struct {
	live    store.SnapshotStore
	archive ArchiveStore // cold / archive tier interface
}

type ArchiveStore interface {
	StoreEnvelope(env interface{}) error // concrete type depends on your envelope package
}

func NewMover(live store.SnapshotStore, archive ArchiveStore) *Mover {
	return &Mover{live: live, archive: archive}
}

func (m *Mover) ArchiveSegment(startID, endID string) error {
	segment, err := m.live.LoadSegment(startID, endID)
	if err != nil {
		return fmt.Errorf("load segment: %w", err)
	}

	for _, env := range segment {
		if err := m.archive.StoreEnvelope(env); err != nil {
			return fmt.Errorf("archive envelope %s: %w", env.ID, err)
		}
	}

	return nil
}
