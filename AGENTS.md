# ILTTU Plugins

This repository is ILTTU's plugin marketplace for Claude Code.
ILTTU itself, its issues and its decisions are in `asteroid-computing/ilttu`, whose `AGENTS.md` holds the rules for how its work is done.

## Project Structure

- `.claude-plugin/marketplace.json` - the marketplace's metadata.
- `plugins/<name>/.claude-plugin/plugin.json` - a plugin's manifest.
- `plugins/ilttu/.mcp.json` - the connection to ILTTU's MCP server.
- `plugins/ilttu/skills/ilttu/SKILL.md` - what ILTTU is, how to check the connection, and what to do when signing in fails.
- `plugins/ilttu/skills/design/` - `/ilttu:design`: designing a page, previewing it as an ILTTU answer and revising it at the same link, with `SKILL.md` and its `references/`.
- `plugins/ilttu/bin/ilttu-design` - the wrapper that downloads `ilttu-design`'s pinned release, checks its SHA-256 and runs it.
- `ilttu-design/` - `ilttu-design`, a Go module on the standard library alone: `serve` previews a design's folder as ILTTU's answers' host serves it, and `init` copies in the design system vendored in `designsystem/`.
- `release-please-config.json` and `.release-please-manifest.json` - the two packages release-please releases, `plugins/ilttu` tagged `ilttu-v<version>` and `ilttu-design` tagged `ilttu-design-v<version>`, and their last versions.
- `.github/workflows/release.yaml` - on main, keeps release-please's release PR open, and when `ilttu-design` releases, builds and publishes its binaries and opens a PR that pins them.
- `.github/workflows/check.yaml` and `pr-title.yaml` - the checks below and the PR title's, on every PR.
- `scripts/pin-ilttu-design` - pins a release's tag and SHA-256s in `plugins/ilttu/bin/ilttu-design`.

## Rules

- **Validate after every change:** `claude plugin validate .` and `claude plugin validate plugins/ilttu`, and in `ilttu-design/`, `go vet ./...` and `go test ./...`.
- **Releases come from release-please, never by hand:** merging its release PR tags each package it releases, so never push a tag or edit a version, a `CHANGELOG.md` or the manifest yourself.
- **A change to `ilttu-design/` reaches people in three merges:** its own PR, the release PR that releases `ilttu-design`, then the pin PR the Release workflow opens, which the next release PR ships in the plugin.
- **`serve` sends what ILTTU's answers' host sends**, checked against `internal/edge` and `internal/cdk/answers` in `asteroid-computing/ilttu`, and its file types are a copy of ILTTU's `internal/filetypes`, so a change there is a change here too.
- **A plugin's change reaches people when its release PR merges**, which changes its `version` in `plugin.json`, because people stay on their copy until the version changes.
  So a change under `plugins/ilttu/` needs a `feat` or `fix` title to be released at all.
- **The connection's values are exact.**
  The URL is `https://mcp.ilttu.com/`, slash included, as Connect on ILTTU's workspace page gives it and as ILTTU's user pool identifies the resource agents sign in for.
  The `oauth` block must stay, because ILTTU's authorization server offers no dynamic client registration, so without a client ID Claude Code can't sign in.
  The callback port is 33418, because Cognito matches a redirect's port exactly.
  The client ID is the SSM parameter `/ilttu-main/cognito/claude-code-client-id` on ILTTU's `main` instance.
- **The repository holds no secret:** a public client's ID, and nothing else from SSM.
- **The install commands in `README.md` are quoted word for word in Connect on ILTTU's workspace page**, so a change to them is a change in `asteroid-computing/ilttu` too.
- **Each skill grows with the MCP server's tools**, and says only what the server does now.

## Writing and pull requests

- One sentence per line in Markdown, never wrapped at a column.
- Voice is ILTTU's: plain words, warm, short.
- Branches are `ajb/<concern>`, and PR titles are Conventional Commits of 72 characters or fewer, because a squash merge uses the title as the commit subject.
