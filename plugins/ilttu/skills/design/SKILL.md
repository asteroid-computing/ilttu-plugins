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
3. **Build the page.** Write `index.html` and what it needs, and look at it in a local browser where there is one, at phone width too.
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

ILTTU serves an answer as it came, on its workspace's own host, with no Content-Security-Policy, so fonts and scripts from other hosts load.

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
