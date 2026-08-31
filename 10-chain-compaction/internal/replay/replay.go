package replay

import (
	"fmt"

	"yourmodule/compaction"
)

// ReplayFromHead walks the chain, verifying compaction snapshots when present.
// Constitutional: identical results whether the segment is live or archived.
func (r *Replayer) ReplayFromHead(headID string) error {
	current := headID

	for current != "" {
		env, err := r.store.GetEnvelope(current)
		if err != nil {
			return fmt.Errorf("load envelope %s: %w", current, err)
		}

		if env.Lineage.CompactionRef != "" {
			snap, err := r.store.GetSnapshot(env.Lineage.CompactionRef)
			if err != nil {
				return fmt.Errorf("load snapshot %s: %w", env.Lineage.CompactionRef, err)
			}

			if err := compaction.VerifySnapshot(snap, r.verifier); err != nil {
				return fmt.Errorf("verify snapshot %s: %w", snap.SegmentStart, err)
			}

			current = snap.SegmentStart
			continue
		}

		if len(env.Lineage.Parents) == 0 {
			break
		}

		current = env.Lineage.Parents[0]
	}

	return nil
}
