# Designing well

> This file is adapted from the frontend-design skill in Anthropic's `anthropics/skills` repository, `skills/frontend-design/SKILL.md` at commit 41bbe19, licensed under the Apache License, Version 2.0 (https://www.apache.org/licenses/LICENSE-2.0).
> It was changed for ILTTU: reworded in ILTTU's voice and reordered, with the design-studio persona left out, and with guidance added on agreeing the look, colour, real material, options, the quality floor, landing pages, phone screens, recreating a design and changing one.

Read this when you plan a look, build a page or a set of options, or when the person pushes back on a design call.
`SKILL.md` has the short version.

## Start from the subject

If the brief doesn't say what the product or subject is, work it out before you design, and check it with the person: one concrete subject, who the design is for, and the one job it does first.
Anything you already know about them or what they're building is a hint.
The subject's world, its materials, its vocabulary and its people, is where distinctive choices come from: a page for a toy for 8-to-11-year-olds looks nothing like a dashboard for financial analysts.
Build with the brief's real content and subject throughout.

## Agree the look

Without a look, references or a design system from the person, get their view before committing to one.
Ask, or show two to four clearly different rough directions they can see, rather than picking a look on their behalf.
A concrete subject, an asset or a brand already counts as their view.
With nobody to ask, commit to one direction you can name, and say which.
Once a look is agreed, it stays agreed.

Where the person has a design system or a brand, use it: its colours, type, spacing, radii and fonts, exactly.
Where it is too tight for what you need, derive from it rather than inventing beside it.

## Plan, check, then build

Work in two passes.

First, write a short plan for the brief, a small token system:

- **Colour:** the base palette as four to six named hex values.
- **Type:** the typefaces and what each is for.
- **Layout:** a layout idea in a sentence or two, with an ASCII sketch to compare ideas, and how content aligns: left, centred or justified.
- **Principles:** what makes this page this page and no other.

Then check the plan against the brief before you build.
If any part reads like the default you'd produce for any similar page, try a similar prompt in your head and see whether you end up in the same place; if you do, change that part, and say what you changed and why.
Only once the plan holds up, write the code, following the revised plan.

For a round of options, make one plan per option, and check each.

When you write the CSS, watch selector specificity: a class for a section and a class for an element can quietly cancel each other out, most often in the padding and margins between sections.

## The generic look

Agent-made design right now clusters around a few looks:

1. a warm cream background, near `#F4F1EA`, with a high-contrast serif display face and a terracotta or clay accent, often near `#D97757`, which is Anthropic's own accent for Claude and so reads as a tell on anyone else's page;
2. a near-black background with one bright acid-green or vermilion accent;
3. a broadsheet layout, with hairline rules, no rounded corners and dense newspaper columns;
4. the SaaS card kit: content cut into identical rounded cards, one radius on everything whatever its rank, the same soft grey shadow, `rgba(0,0,0,.1)`, under each, and gradient washes as decoration;
5. template chrome that turns up whatever the subject: a tracked-out all-caps eyebrow over every heading, meta joined with middle dots (`A · B · C`), labels built as "WORD — fragment", near-black like `#0B0B0B` or `#111` standing in for black, a monospace face for small data labels, and a `→` on the end of link and button text.

Watch for these typographic tells too:

- one word or phrase in a headline picked out in italic, bold or another colour;
- all caps for labels;
- labels above content that don't need to be there.

And these tropes:

- gradient backgrounds used heavily;
- emoji, unless the brand uses them;
- rounded containers with a coloured left border.

Each of these suits some brief, but they are defaults, not choices, and they appear whatever the subject.
Where the brief pins down a look, follow it exactly, even when it asks for one of these: the brief's own words always win.
Where it leaves something open, don't spend that freedom on a default.

## Type

Type carries the page's personality.
One family or two is plenty, and if two, make them clearly different.
Choose typefaces for this brief, not the families you'd reach for on any project, and set a clear type scale with deliberate weights, widths and spacing, in the spirit of The Elements of Typographic Style.
When type is a headline or a visual element, make the treatment itself part of the design, not a neutral carrier for the words.
Keep lines under about 80 characters; serif text can run a little longer, and serif body text wants a little more line height than sans.

## Colour

Choose a tone for the page, warm, cool or neutral, and tint its whites and blacks a touch toward it rather than using pure `#fff` and `#000`.
Use no more than two accents, written in `oklch`, with like chroma and lightness and different hues, so they sit together.
Take the brand's colours first, and derive from them before inventing new ones.

## Layout and structure

The hero is the first thing anyone sees on a web page.
Open with the most characteristic thing in the subject's world, in whatever form suits it: a headline, an image, an animation, a live demo or something to try.
A big number with a small label, a few supporting stats and a gradient accent is the default hero, so use it only when it's truly the best choice.

Structure is information.
Outlines, borders, numbers, eyebrows, dividers and labels should say something about the content, not decorate it.
Numbered markers like 01, 02, 03 are for content that really is a sequence, such as steps or a timeline: check before you add them.

## Motion

Use motion nobody asked for sparingly, and only to draw the eye.
One orchestrated moment, a single load sequence or one reveal, lands better than effects scattered everywhere; a fade-and-slide on every section and a hover effect on every card are the default, and read as generated.
Motion that answers what someone did, opening, expanding, confirming, is welcome when it shows what changed.

## Real material

A detailed mockup starts from the person's own material: their code, their brand, screenshots of their product, their design system.
Find that first, and ask for it if you can't.
Mocking a whole product from nothing is the last resort, and it shows.
Say your assumptions and reasoning early, and share something as soon as there is something to react to.

Where you don't have an icon, an image or a component, draw a clear placeholder: a placeholder beats a poor imitation of the real thing.

## Options

When a big decision is still open, such as the overall direction, the hero, the type, the colour or the density, settle it before you build the full page.

- Offer two to four clearly different options, each on an axis you can name, such as "warm editorial" against "dense, data first"; five shades of one look is no choice at all.
- Rough is fine for choosing a direction: early on, sketch three to five approaches with simple shapes, placeholder text and little colour, focused on structure and flow.
- Make each option's honest case, and its main tradeoff, in your reply, not on its page; a set where only your favourite gets a case is a rigged vote.
- Keep option names stable: once something is "Option B" or "Warm editorial", it keeps that name in every later round.
- Once the direction is settled and the person wants variations to keep, give three or more across several dimensions, from by-the-book versions to new layouts, interactions and visual ideas, so they can mix and match.

## The floor

Build to a quality floor without announcing it:

- It works at phone width, with no headline breaking badly, no squashed grid and no text too small to read.
- Keyboard focus is visible.
- Reduced motion is respected.
- Touch targets are at least 44px, and text is sized for how far away it's read.
- Controls are the real elements even in a static mockup: `<button>`, `<a href>`, and `<input>` with a `<label>`, never a `div` with a click handler, which Tab skips and a screen reader can't announce.
- An icon-only button has an `aria-label`.
- Text has a contrast of at least 4.5:1 against what's behind it, or 3:1 from 24px; small grey captions and white text on coloured fills fail most often, so darken them.
- Colours that must be told apart differ in lightness, not hue alone; blue and orange separate better than red and green.
- The palette is harmonious.

## Landing pages

A landing page has a known anatomy:

- a hero that states the offer in one sentence, with one clear call to action;
- proof a visitor can trust, such as quotes, logos or numbers, from the person's own material or as clearly marked placeholders;
- sections that answer a visitor's real doubts rather than listing features;
- one main action, repeated down the page, not three buttons competing.

On a landing page the copy is the product.
Write specific copy from what the person told you, about their product, their customers and in their voice: never lorem ipsum, "Welcome to our website" or filler that could describe any business.
Where a fact is missing, such as a price, a date or an address, put a marked placeholder like `[YOUR PRICE]` for them to fill, never a made-up one.
It's a page, not a poster: fluid to the window's width, down to a phone.

## Phone screens

Don't draw a fake status bar, the strip with the time, battery and signal, or a fake keyboard.
On a real phone the real ones sit on top of your layout, and a painted copy looks doubled.
Leave that space alone, inside a drawn phone frame too.

## Recreating a design

When asked to recreate a design whose source you can reach, such as a repository, pasted files or a design system, build from that source, not from what you remember of the product.
Read its components and styles, use the assets it actually loads, and copy its exact values, paddings, radii, font sizes and line heights, rather than rounding them to a grid or a framework's defaults.
Source beats screenshots: with the source in hand, treat screenshots as a rough guide.
If you can't read the source, say so rather than inventing from memory.

Don't recreate another company's distinctive design, its proprietary patterns or its branded visuals, unless the person says they work there; ask when it's unclear.
Instead, find out what they want to build, and make an original design that does it.

## Changing a design

- **A small change stays small.** When asked to change some text, a colour or one element, change only that, and leave every other layout, spacing, font, size, colour and word as it was. A redesign or a new direction is different, and then the big changes are the job. If a broader change would help, make the small one and suggest the rest.
- **Follow the design's own style.** Before adding to an existing page, learn how it does things, its copy, palette, tone, states, motion, cards, shadows and density, and do the same.
- **No filler.** Never pad a page with placeholder text, extra sections or numbers just to fill space; every element earns its place. An empty-feeling section is a layout problem, not a content one. Leave out stats, icons and data nobody needs.
- **Ask before adding.** If more sections, pages or copy would help, ask first: the person knows their audience and goals better than you do.

## Words

Words are in a design for one reason, to make it easier to understand and use.
They're design content, not decoration: ask what the page needs to say, and how to say it so it helps someone find their way.

- Write from the reader's side, naming things by what they'll understand, not by how the system is built: someone manages notifications, not webhook config.
- Say what something is or does in plain words rather than selling it; specific and clear beats clever.
- Use the active voice, and let a button say exactly what happens: "Save changes", not "Submit".
- Keep an action's name through the whole flow, so "Publish" leads to "Published".
- Treat errors and empty screens as direction: say what went wrong and how to fix it, without apologising or being vague, and make an empty screen an invitation to act.
- Keep it conversational: plain verbs, sentence case, no filler, and a tone matched to the brand and the reader, with each piece of text doing one job.

## Look at your work

Spend your boldness in one place: let one element be the memorable thing, keep everything around it quiet, and cut decoration the brief doesn't need.
As you build, look at the page in screenshots where you can, at desktop and phone width; a picture says more than reading the code back.
Before you hand it over, take one thing away, as Chanel said to do with accessories before leaving the house.
Where you can keep notes between rounds, write down what you've tried, so the next round tries something new.
