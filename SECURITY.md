# Security

## Reporting a Vulnerability

Please report suspected vulnerabilities privately to `aj@asteroidcomputing.com`.
Include the affected plugin, a short impact summary, and reproduction steps when you can share them safely.

## Credentials

This repository holds no secret.
The client ID in `plugins/ilttu/.mcp.json` belongs to a public OAuth client, which has no client secret and signs in with PKCE, so it is safe to publish.
The tokens from signing in are Claude Code's to store and refresh, and they never reach this repository.

Do not commit tokens, workspace keys, local tool settings or generated authentication files to this repository.

## Downloads

`plugins/ilttu/bin/ilttu-design` downloads one release of `ilttu-design` from this repository's GitHub Releases, over HTTPS, the first time it runs, into the plugin's data folder or the user's cache.
It runs the binary only while its SHA-256 is the one pinned in the wrapper for that platform, and refuses it otherwise.
A release is built by `.github/workflows/release.yaml` from the tag release-please makes on `main`, and published only once its binaries and `SHA256SUMS` are attached.
The wrapper's pins change only through a reviewed PR, which the same workflow opens from the published `SHA256SUMS`.
