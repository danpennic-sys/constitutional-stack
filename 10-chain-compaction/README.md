# 10 – Chain Compaction Semantics

Frozen constitutional layer for replacing a contiguous prefix of an envelope chain with a single verifiable digest while preserving full replay and lineage.

## Contents

- `internal/compaction/snapshot.go` – Snapshot type, root-hash, digest, create & verify
- `internal/envelope/lineage.go` – `compaction_ref` extension
- `internal/replay/replay.go` – Compaction-aware replay loop
- `server/api/snapshots.post.ts` / `[id].get.ts` – Nitro routes
- `cmd/nitro-compact/main.go` – Operator CLI
- `migrations/001_compaction.sql` – PostgreSQL schema
- `internal/compaction/conformance_test.go` – CP-01 … CP-08 skeletons

## Status

Implementation complete and ready for integration.
