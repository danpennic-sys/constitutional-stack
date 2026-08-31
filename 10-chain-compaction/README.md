# 10 — Chain Compaction Semantics

![Status](https://img.shields.io/badge/status-implemented-brightgreen)
![CP Tests](https://img.shields.io/badge/CP--01→08-defined-blue)

> Replace a contiguous prefix of an envelope chain with a single verifiable digest  
> while preserving full lineage, replay, and proof binding.

---

## Package Contents

| Path | Purpose |
|------|---------|
| `internal/compaction/snapshot.go` | Snapshot type, root-hash, digest, create & verify |
| `internal/store/snapshot_store.go` | Constitutional store interface |
| `internal/store/postgres_snapshots.go` | Postgres implementation |
| `internal/archive/mover.go` | Archive mover (relocate only) |
| `internal/envelope/lineage.go` | `compaction_ref` extension |
| `internal/replay/replay.go` | Compaction-aware replay loop |
| `server/api/snapshots/` | Nitro routes |
| `cmd/nitro-compact/` | Operator CLI |
| `migrations/` | Frozen schema + compaction_ref |
| `internal/compaction/conformance_test.go` | CP-01 → CP-08 skeletons |

---

## Constitutional Guarantees

- Compaction is **append-only** — envelopes are never deleted
- Snapshot fields are immutable and explicit
- `created_at` is persistence metadata only (never canonical)
- Signature stored as raw `BYTEA`
- Replay with snapshots produces identical results to full replay
- Old verifiers continue to work on non-compacted chains

---

## Quick Operator Flow

```bash
# Dry run
nitro-compact --start <uuid> --end <uuid> --dry-run

# Live (after governance approval)
nitro-compact --start <uuid> --end <uuid>
```

---

## Conformance Vectors

- **CP-01** Correct segment hashing
- **CP-02** Signature verifies
- **CP-03** Root hash matches archive
- **CP-04** Chain traversable with `compaction_ref`
- **CP-05** Replay identical pre/post compaction
- **CP-06** Old verifier still works
- **CP-07** Mixed chains behave deterministically
- **CP-08** Proofs remain bound to envelopes
