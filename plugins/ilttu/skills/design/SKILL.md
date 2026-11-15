---
name: design
description: Designs a page, a screen or a set of options and shares it as a private ILTTU answer to preview, then revises it at the same link. Use when someone wants to show, share or preview a design through ILTTU, or to change a design already published as an ILTTU answer.
---
# Design

This skill designs a page, shares it as an ILTTU answer the person opens at a private link, and revises it at that same link.
It works through ILTTU's MCP server, which this plugin connects; if its tools aren't there, the `ilttu` skill says how to sign in.

## The loop

1. **Settle the brief.** Read what the person gave you, and ask only what would change the design, in one round: `references/questions.md`.
2. **Choose the look.** Plan the palette, type and layout, and check the plan against the generic look before you build: `references/craft.md`.
3. **Build and preview the page.** Write `index.html` and what it needs in a folder of its own, and preview it with `ilttu-design serve`: see "Preview it" below.
4. **Publish it.** Upload it with `start_upload`, then read it with `get_answer` until it's ready: `references/publishing.md`.
5. **Hand back the link.** Say what it shows, who can open it, and what you'd like them to tell you.
6. **Revise it.** Change what they ask for, and replace the answer with `replace_answer`, so the link stays the same.

## The page's shape

- One `index.html` at the top, with its styles, scripts, images and fonts beside it, linked by relative paths.
- A page that is one file goes up as a `.html` file; one with other files goes up as a `.zip`, `.tar`, `.tar.gz` or `.tgz`.
- Fluid, not fixed: a `max-width` container with side padding, and grids that fall to one column at phone width.
- Fonts from Google Fonts, as a `<link>` in the `<head>`, each with a system fallback.
- Scripts from a CDN pinned to an exact version, such as `htmx.org@2.0.4`, never a range or `latest`.
- A `<title>`, which becomes the answer's title.

ILTTU serves an answer as it came, with no Content-Security-Policy, so fonts and scripts from other hosts load.

## Preview it

`ilttu-design` serves a design's folder on 127.0.0.1 as ILTTU will serve it as an answer: with the answer's headers, no Content-Security-Policy, and `no-store`, so a reload always shows the latest.
It answers a 404 that says why for any file ILTTU would refuse or never serve, such as one with no extension, a symbolic link, or a path under `/_/`.
The plugin puts it on the Bash tool's PATH, and it downloads its pinned release the first time it runs.

- `ilttu-design serve <dir>` serves the folder at a free port and prints its URL: run it in the background.
- `ilttu-design init <dir>` copies the ilttu design system into `<dir>/ilttu/`: `tokens.css`, `components.css`, and its fonts with their licences.
  Use it only for a design in ILTTU's own look, link `ilttu/tokens.css` and then `ilttu/components.css` in each page's `<head>`, and run it again to update them; it replaces only its own files.

In the desktop app, preview it in the built-in browser with `preview_start`, from this `.claude/launch.json` at the project's top, with the design's folder, relative to the project's top, in `runtimeArgs`:

```json
{
  "version": "0.0.1",
  "configurations": [
    {"name": "design", "runtimeExecutable": "${CLAUDE_PLUGIN_ROOT}/bin/ilttu-design", "runtimeArgs": ["serve", "design"], "port": 8724, "autoPort": true}
  ]
}
```

The path in `runtimeExecutable` changes when the plugin updates, so write it as this skill gives it, and rewrite it when the file's differs.
Elsewhere, open the URL `serve` prints in a headless browser.
Look at the page at 1,440 and 375 wide, in light and dark, before you hand it back.

## Rules that bite

- A design is an HTML answer, never Markdown: `publish_answer` and `.md` files render text in one of ILTTU's styles, not your page.
- A new answer takes its workspace's default for who can open it: say which it is, and never make it open to anyone with the link unless the person asks.
- Replacing an answer removes what it served before, for good, so say what will go before you replace it.
- On any refusal from `replace_answer`, read the answer again with `get_answer` and go by what it says, never by the refusal's words.

## Designing well

The full guidance is in `references/craft.md`.

- Agree the look with the person before you build, or show them two to four clearly different rough directions.
- Plan before you build, and check the plan against the generic look an agent reaches for on any page.
- Build from their real material; where a fact is missing, put a marked placeholder like `[YOUR PRICE]`, never an invented one.
- No filler, and ask before adding anything they didn't ask for.
- A small change stays small.
- Accessible as drawn: real buttons, links and labels, visible focus, text at 4.5:1, and reduced motion respected.
- Look at your own work in screenshots before you hand it back.
- Each option's case goes in your reply, never on the page.

## References

- `references/craft.md`: designing well, from the brief to checking your own work.
- `references/questions.md`: when and how to ask before you build.
- `references/publishing.md`: packing the page, uploading and replacing it, options, who can open it, and the hand-back.
