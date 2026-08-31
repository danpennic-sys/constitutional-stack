package replay

import (
	"fmt"

	"yourmodule/compaction"
)

func (r *Replayer) ReplayFromHead(headID string) error {
	current := headID
	for current != "" {
		env, err := r.store.GetEnvelope(current)
		if err != nil {
			return err
		}

		if env.Lineage.CompactionRef != "" {
			snap, err := r.store.GetSnapshot(env.Lineage.CompactionRef)
			if err != nil {
				return fmt.Errorf("snapshot fetch: %w", err)
			}
			if err := compaction.VerifySnapshot(snap, r.verifier); err != nil {
				return fmt.Errorf("snapshot invalid: %w", err)
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
