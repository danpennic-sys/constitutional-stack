# Constitutional Stack

Umbrella repository for the complete frozen constitutional contract.

## License

Apache License 2.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE).

**Dual-use / cryptographic notice:** This software includes cryptographic
functionality. Users are responsible for compliance with applicable export
control and dual-use regulations. See the NOTICE file for the full statement.

## Layers

| # | Layer | Directory | Status |
|---|-------|-----------|--------|
| 1 | Envelope contract | `01-envelope-contract/` | Skeleton |
| 2 | Lineage schema | `02-lineage-schema/` | Skeleton |
| 3 | Evolution rules | `03-evolution-rules/` | Skeleton |
| 4 | Verifier identity + key rotation | `04-verifier-identity/` | Skeleton |
| 5 | Proof schema | `05-proof-schema/` | Skeleton |
| 6 | Evidence attachment | `06-evidence-attachment/` | Skeleton |
| 7 | Replay protocol | `07-replay-protocol/` | Skeleton |
| 8 | Chain query semantics | `08-chain-query/` | Skeleton |
| 9 | Verifier upgrade protocol | `09-verifier-upgrade/` | Skeleton |
| 10 | Chain compaction semantics | `10-chain-compaction/` | **Implemented** |

## Compaction (Layer 10)

Fully implemented Go snapshot core, Nitro API routes, CLI, migrations, and conformance test skeletons.
See `10-chain-compaction/README.md`.

## Next steps

- Flesh out the remaining layer directories with the corresponding frozen contracts and implementations.
- Keep each layer independently versionable while remaining under this single constitutional umbrella.
