# ILTTU Plugins

This repository is ILTTU's plugin marketplace for Claude Code.
ILTTU itself, its issues and its decisions are in `asteroid-computing/ilttu`, whose `AGENTS.md` holds the rules for how its work is done.

## Project Structure

- `.claude-plugin/marketplace.json` - the marketplace's metadata.
- `plugins/<name>/.claude-plugin/plugin.json` - a plugin's manifest.
- `plugins/ilttu/.mcp.json` - the connection to ILTTU's MCP server.
- `plugins/ilttu/skills/ilttu/SKILL.md` - what ILTTU is, how to check the connection, and what to do when signing in fails.

## Rules

- **Validate after every change:** `claude plugin validate .` and `claude plugin validate plugins/ilttu`.
- **Change a plugin's `version` in `plugin.json` with every change to it**, because people stay on their copy until the version changes.
- **The connection's values are exact.**
  The URL is `https://mcp.ilttu.com/`, as Connect on ILTTU's workspace page gives it.
  The `oauth` block must stay, because ILTTU's authorization server offers no dynamic client registration, so without a client ID Claude Code can't sign in.
  The callback port is 33418, because Cognito matches a redirect's port exactly.
  The client ID is the SSM parameter `/ilttu-main/cognito/claude-code-client-id` on ILTTU's `main` instance.
- **The repository holds no secret:** a public client's ID, and nothing else from SSM.
- **The install commands in `README.md` are quoted word for word in Connect on ILTTU's workspace page**, so a change to them is a change in `asteroid-computing/ilttu` too.
- **The skill grows with the MCP server's tools**, and says only what the server does now.

## Writing and pull requests

- One sentence per line in Markdown, never wrapped at a column.
- Voice is ILTTU's: plain words, warm, short.
- Branches are `ajb/<concern>`, and PR titles are Conventional Commits of 72 characters or fewer, because a squash merge uses the title as the commit subject.
