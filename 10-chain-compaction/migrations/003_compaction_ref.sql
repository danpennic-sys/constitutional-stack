-- Add compaction_ref to envelopes (nullable FK)
ALTER TABLE envelopes
    ADD COLUMN IF NOT EXISTS compaction_ref UUID NULL REFERENCES snapshots(snapshot_id);
