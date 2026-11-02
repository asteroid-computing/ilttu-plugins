# Security

## Reporting a Vulnerability

Please report suspected vulnerabilities privately to `aj@asteroidcomputing.com`.
Include the affected plugin, a short impact summary, and reproduction steps when you can share them safely.

## Credentials

This repository holds no secret.
The client ID in `plugins/ilttu/.mcp.json` belongs to a public OAuth client, which has no client secret and signs in with PKCE, so it is safe to publish.
The tokens from signing in are Claude Code's to store and refresh, and they never reach this repository.

Do not commit tokens, workspace keys, local tool settings or generated authentication files to this repository.
