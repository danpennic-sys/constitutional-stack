# Constitutional Stack

![Status](https://img.shields.io/badge/status-constitutional-blueviolet)
![Layers](https://img.shields.io/badge/layers-10-blue)
![License](https://img.shields.io/badge/license-Apache%202.0-green)
![Crypto](https://img.shields.io/badge/crypto-dual--use%20notice-orange)

> Frozen constitutional contract for verifiable envelope chains.
> Deterministic. Replay-safe. No inference. No drift.

---

## The Ten Layers

| # | Layer | Directory | Status |
|---|-------|-----------|--------|
| 01 | Envelope Contract | [`01-envelope-contract/`](01-envelope-contract/) | Skeleton |
| 02 | Lineage Schema | [`02-lineage-schema/`](02-lineage-schema/) | Skeleton |
| 03 | Evolution Rules | [`03-evolution-rules/`](03-evolution-rules/) | Skeleton |
| 04 | Verifier Identity + Key Rotation | [`04-verifier-identity/`](04-verifier-identity/) | Skeleton |
| 05 | Proof Schema | [`05-proof-schema/`](05-proof-schema/) | Skeleton |
| 06 | Evidence Attachment | [`06-evidence-attachment/`](06-evidence-attachment/) | Skeleton |
| 07 | Replay Protocol | [`07-replay-protocol/`](07-replay-protocol/) | Skeleton |
| 08 | Chain Query Semantics | [`08-chain-query/`](08-chain-query/) | Skeleton |
| 09 | Verifier Upgrade Protocol | [`09-verifier-upgrade/`](09-verifier-upgrade/) | Skeleton |
| **10** | **Chain Compaction Semantics** | [`10-chain-compaction/`](10-chain-compaction/) | **Implemented** |

---

## Layer 10 — Compaction (Live)

Full operator-grade implementation:

- Snapshot generator & verifier (Go)
- Constitutional Postgres schema
- Snapshot store interface + Postgres implementation
- Archive mover (relocate, never delete)
- Nitro API routes
- CLI (`nitro-compact`)
- Conformance vectors CP-01 → CP-08

See [`10-chain-compaction/README.md`](10-chain-compaction/README.md).

---

## License & Dual-Use Notice

Apache License 2.0 — see [LICENSE](LICENSE).

This software contains **cryptographic functionality**.  
Users are responsible for compliance with applicable export-control and dual-use regulations.  
Full statement in [NOTICE](NOTICE).

---

## Constitutional Principles

- Append-only
- Deterministic canonical bytes
- No inference of identity or lineage
- Replay must produce identical results with or without compaction
- Proofs remain bound to envelopes forever

---

*Operator Pages · Constitutional Stack*
