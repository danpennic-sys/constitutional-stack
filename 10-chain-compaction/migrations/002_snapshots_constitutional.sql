-- Constitutional snapshot storage
-- Immutable fields only. No inference. No persistence metadata in canonical identity.

CREATE TABLE IF NOT EXISTS snapshots (
    snapshot_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    type TEXT NOT NULL CHECK (type = 'compaction_snapshot'),

    segment_start UUID NOT NULL REFERENCES envelopes(event_id),
    segment_end   UUID NOT NULL REFERENCES envelopes(event_id),

    root_hash TEXT NOT NULL,
    compaction_digest TEXT NOT NULL,

    verifier_id TEXT NOT NULL,
    version TEXT NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,

    signature BYTEA NOT NULL,

    -- Persistence metadata only — never enters canonical bytes
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_snapshots_segment_start ON snapshots(segment_start);
CREATE INDEX IF NOT EXISTS idx_snapshots_segment_end   ON snapshots(segment_end);
