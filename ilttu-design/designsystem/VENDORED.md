# The design system, vendored

These files are the template assets' one source: `ilttu-design init` writes `tokens.css` from `tokens.json`, and copies `bundle.css`, as `components.css`, and the fonts with their licences as they are.

They are copied from the "ilttu" Design System Artifact, https://claude.ai/artifact/YS1gEZBcZyEnwuWvvUXEQE, at version `1794733012-49f8`, last changed 2026-11-15 (#711).
The Artifact is private and nothing can fetch it at build time, so this copy is pinned here.
The fonts are the same files as ILTTU's own `internal/web/assets/fonts/`, whose `VENDORED.md` names where each came from, and each licence is copied from there.

| File                                           | From                                 | SHA-256                                                            |
| ---------------------------------------------- | ------------------------------------ | ------------------------------------------------------------------ |
| `tokens.json`                                  | `project/tokens.json`                | `9be8be3aaaea0702e739bfbbcb0fdab816fcbe50c3470343343a39c3ad2161a1` |
| `bundle.css`                                   | `project/components/bundle.css`      | `42cba781466b9546fa76cb3faba734d798000925d2836bacd065d7bc9f7048c2` |
| `fonts/gabarito-latin-wght-normal.woff2`       | `project/fonts/`                     | `780dd129f53815f23c37d4b786c98139c3cfa3086c391ec0bb9aa40d26e7a8fe` |
| `fonts/figtree-latin-wght-normal.woff2`        | `project/fonts/`                     | `4ba7d3d096695818fe0686be4f1e82c6b05134e18a22260336130335027462dd` |
| `fonts/jetbrains-mono-latin-wght-normal.woff2` | `project/fonts/`                     | `18be452724bfdc236c074ca94a249a7f41a86752c7d04ab258ce9ed5651f6a7e` |
| `fonts/OFL-gabarito.txt`                       | ILTTU's `internal/web/assets/fonts/` | `c51586fc532deab05c1aa9da0fd240ae58af6967e505d3f7f6562476f17f7a1d` |
| `fonts/OFL-figtree.txt`                        | ILTTU's `internal/web/assets/fonts/` | `ee23e6c84000126692e112ff067b456470493349f993ceaa8b4d3766dd6fad5d` |
| `fonts/OFL-jetbrains-mono.txt`                 | ILTTU's `internal/web/assets/fonts/` | `403581b69dac5cff4079205e01c6b467e56af449ecbd7247693ddb1baafa005b` |

To update them, read the Artifact's files again, replace these, update the version and each row, and cut a release.
