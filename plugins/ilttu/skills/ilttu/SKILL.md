---
name: ilttu
description: What ILTTU is, how to check the connection to it with whoami, and what to do when signing in to it fails. Use when someone asks about ILTTU or their connection to it, or when ILTTU's MCP server needs authentication or answers with an error.
---
# ILTTU

ILTTU ("I'd love to tell you") lets agents publish HTML and hand back a private link.
What it publishes are answers, and answers live in workspaces.
This plugin connects Claude Code to ILTTU's MCP server, `https://mcp.ilttu.com/`, which acts for the person who signed in.
The server's tools grow over time, so go by the tools it lists now.

## Check the connection

Call ILTTU's `whoami` tool, and tell the person who it says they are.

- `kind` `person`: the connection acts for a person, and `email` says who.
- `kind` `key`: the connection acts for a workspace key, and `workspace`, `name` and `role` say which.

If the tool isn't there, the server needs signing in.

## Sign in

ILTTU is invite-only, so the person needs an invitation first.

- **In a terminal session:** run `/mcp`, choose ILTTU's server and choose Authenticate.
- **In the desktop app:** Reconnect doesn't start a sign-in, so run `claude mcp login plugin:ilttu:ilttu` in a shell.

The browser opens ILTTU's sign-in page on `auth.ilttu.com`, then returns to Claude Code.
Never ask for the token, and never put it in a message or a file.

## When signing in fails

| What happens                                          | What it means                                                                                         | What to do                                                                                                                                                                              |
| :---------------------------------------------------- | :---------------------------------------------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `OAuth callback port 33418 is already in use`         | Another sign-in, often in another session, holds the port                                             | Finish or cancel that sign-in, then try again; `lsof -ti:33418 -sTCP:LISTEN` finds what holds it                                                                                        |
| After signing in, the browser can't reach `localhost` | The redirect didn't reach Claude Code                                                                 | Paste the full URL from the browser's address bar into Claude Code's prompt                                                                                                             |
| The sign-in page refuses the redirect                 | Claude Code sent a callback other than `http://localhost:33418/callback`                              | Update Claude Code, then try again                                                                                                                                                      |
| Calls answer 403                                      | The person has never signed in to ILTTU in a browser                                                  | Sign in at `https://ilttu.com` once, then try again                                                                                                                                     |
| Calls answer 404                                      | The person has never signed in to ILTTU in a browser, or ILTTU's agents feature isn't on for them yet | Sign in at `https://ilttu.com` once and try again; if it still answers 404, ask whoever invited them                                                                                    |
| Sign-in completes, then every call answers 401        | ILTTU doesn't accept the token                                                                        | Sign in again once; if it still answers 401, open an issue at `https://github.com/asteroid-computing/ilttu-plugins/issues` with the response's `WWW-Authenticate` line, never the token |
| The server shows as `ilttu`, not `plugin:ilttu:ilttu` | It was added by hand before, and Claude Code keeps that one over the plugin's for the same URL        | It reaches the same server, so nothing needs fixing; `claude mcp remove ilttu` hands over to the plugin's                                                                               |
