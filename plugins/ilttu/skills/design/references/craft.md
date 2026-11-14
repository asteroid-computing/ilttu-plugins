# Designing well

> This file is adapted from the frontend-design skill in Anthropic's `anthropics/skills` repository, `skills/frontend-design/SKILL.md` at commit 41bbe19, licensed under the Apache License, Version 2.0 (https://www.apache.org/licenses/LICENSE-2.0).
> It was changed for ILTTU: reworded in ILTTU's voice and reordered, with the design-studio persona left out, and with guidance added on agreeing the look, colour, real material, options, the quality floor, landing pages, phone screens, recreating a design and changing one.

Read this when you plan a look, build a page or a set of options, or when the person disagrees with a choice you made.
`SKILL.md` has the short version.

## Start from the subject

If the brief doesn't say what the product or subject is, work it out before you design, and check it with the person: one concrete subject, who the design is for, and the one job it does first.
Anything you already know about them or what they're building is a hint.
The subject's world, its materials, its vocabulary and its people, is where distinctive choices come from: a page for a toy for 8-to-11-year-olds looks nothing like a dashboard for financial analysts.
Build with the brief's real content and subject throughout.

## Agree the look

The look is the person's call, so find out what they want before you settle it.
A brand, an asset, references or even a clear subject is enough to go on.
If they gave you none of those, ask, or put a few rough pages in front of them that differ in an obvious way and let them point.
If nobody is there to answer, pick one look, give it a name, and tell them which you picked.
Once a look is agreed, it stays agreed.

Where the person has a design system or a brand, use it: its colours, typefaces, spacing and corner radii, exactly.
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

Good detail comes from what already exists.
Before you draw someone's product, look for its real pieces, and ask when you can't find them.
A product invented from nothing looks invented.
Tell the person early what you assumed and why, and show them a first version while it's still easy to change.
When a real icon, picture or component is missing, leave a plain, labelled gap where it goes rather than faking one.

## Options

When a big decision is still open, such as the overall direction, the hero, the type, the colour or the density, settle it before you build the full page.

- Show two to four, and make each differ in a way you can say out loud, like "a calm reading page" against "everything on one screen"; if two only differ in shade, drop one.
- Keep them rough while the person is choosing: boxes, grey text and one colour are enough to compare how a page is put together.
- In your reply, argue fairly for every option and say what each one costs, and keep the pitch off the pages themselves.
- Name each option and keep the name, so what's called "Option B" today is still "Option B" next round.
- When the direction is set and the person wants alternatives to keep, vary more than one thing, some close to what exists and some trying a new layout, interaction or visual idea, so they can take a piece from each.

## The floor

Build to a quality floor without announcing it:

- It works at phone width, with no headline breaking badly, no squashed grid and every word big enough to read.
- Keyboard focus is visible.
- Reduced motion is respected.
- Touch targets are at least 44px, and text is sized for how far away it's read.
- Buttons, links and form fields are the real HTML elements, each field with its label, even when nothing on the page works yet, so a keyboard and a screen reader can find them.
- A button that shows only an icon still has a name, in `aria-label`.
- Text is easy to read against its background, at the 4.5:1 contrast WCAG asks for, or 3:1 for large text; pale captions and white on a bright fill are where it usually slips.
- Where colour carries meaning, the colours differ in how light they are as well as in hue, so someone who can't tell hues apart still can.
- The palette is harmonious.

## Landing pages

A landing page exists to get a visitor to do one thing.

- The first screen says what's on offer and what to do about it.
- Somewhere near it, the page shows why a visitor should believe it, with the person's own quotes, customers or figures, or a marked gap for them to fill.
- Each section after that answers a question a visitor would really ask, rather than walking through a feature list.
- The one action comes back further down the page, and nothing competes with it.

Here the words matter as much as the look.
Write them from what the person told you about their product and the people who buy it, in a voice that sounds like them, because a line that would fit any business tells a visitor nothing.
If you don't know a price, a date or an address, leave a clear marker such as `[price]`, and never guess one.
Build it to stretch and shrink with the window, from a wide screen down to a phone, rather than at one fixed size.

## Phone screens

Leave out the phone's own parts, the clock and battery along the top and the on-screen keyboard.
The phone draws its own over your page, so a drawn copy ends up on screen twice.
That holds inside a phone you draw as a frame, too.

## Recreating a design

To rebuild something that exists, work from its source if you can reach it, whether that's a repository, files they pasted or a design system.
Use what's really there, its components, its styles and the assets it loads, and keep its measurements as they are, without tidying them to a grid.
Memory is a poor guide to how a product looks, and once you have the source, a screenshot is only a rough one.
If you can't get at the source, tell the person instead of building from memory.

Another company's signature look isn't yours to copy, unless the person tells you they work there, so ask if you're not sure.
Find out what they're trying to make, and design something of their own that does it.

## Changing a design

- **Do what was asked, and only that.**
  When someone wants new text, a new colour or one element changed, everything else stays exactly as it was.
  A redesign is the opposite case, and then large changes are the point.
  If you see a bigger improvement, make the small change and mention the bigger one.
- **Match what's there.**
  Before you add to a page, notice how it already writes, colours, spaces, moves and shades things, and add in the same way.
- **Leave space empty rather than filling it.**
  Don't add text, sections or figures to make a page look full, and if a section looks bare, fix the layout.
  Every number, icon and chart is one somebody needs.
- **Ask before you add.**
  More pages, sections or copy are the person's call, since they know who it's for.

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
