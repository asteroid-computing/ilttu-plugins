# ilttu-plugins

ILTTU's plugin marketplace for Claude Code.
ILTTU ("I'd love to tell you") lets agents publish HTML and hand back a private link.

## Add the plugin

In Claude Code:

```shell
/plugin marketplace add asteroid-computing/ilttu-plugins
/plugin install ilttu@ilttu-plugins
```

Or from a shell:

```shell
claude plugin marketplace add asteroid-computing/ilttu-plugins
claude plugin install ilttu@ilttu-plugins
```

Then sign in, with an invitation to ILTTU: run `/mcp` and choose ILTTU's server, or in the desktop app run this in a shell.

```shell
claude mcp login plugin:ilttu:ilttu
```

Ask Claude to call ILTTU's `whoami`, and it answers with who you signed in as.

To get a newer version, run `/plugin marketplace update ilttu-plugins`.

## Plugins

### `ilttu`

- It connects Claude Code to ILTTU's MCP server, `https://mcp.ilttu.com/`, signing you in through ILTTU's own page.
- The `ilttu` skill says what ILTTU is, how to check the connection, and what to do when signing in fails.

This plugin is for `ilttu.com`.
For another instance of ILTTU, use the command in Connect on a workspace's page.

## Validate

```shell
claude plugin validate .
claude plugin validate plugins/ilttu
```

See [SECURITY.md](./SECURITY.md) for how sign-in and credentials work and how to report a vulnerability.

## License

[Apache-2.0](./LICENSE).
