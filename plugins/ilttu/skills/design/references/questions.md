# Asking before you build

Read this before you ask the person anything about a design.

## Whether to ask

Ask only when someone is there to answer.
In a run with nobody to ask, such as a scheduled task or another agent calling you, or when the person says "just make it", don't ask: decide, build, and say your assumptions in one line.

Read everything they gave you first.
Ask when what they gave still leaves two or more decisions open that would change what you build.
A point you can sensibly default, a full brief or a small revision needs no questions.

## How to ask

Ask once, before you write anything, with at most four questions.
Where your tools have one, such as Claude Code's `AskUserQuestion`, use it, so they can answer with a click.

Write each question from their material, not from a form.
Name what you read, the gap or tension in it, and the decision it leaves open, with the decision that would most change the design first.
Never ask what the conversation already answers.
Default what the setting implies, and say so.
Stock questions, about audience, length or how polished it should be, are only for a brief with almost nothing in it.

The options do the real work:

- Two to four concrete choices for this design, such as "lead with the booking form", not "functional".
- Each differs on something you can name, never shades of one idea.
- Labels of a few words that carry the choice, with a short description where it helps.
- Your pick first, marked "(Recommended)", except on a question of fact.
- Let them pick more than one where the options don't exclude each other.
- No "Other" or "you decide" option, as the tool adds its own, and no question that only free text can answer.

## After they answer

Treat the answers as decisions, and restate them in your one line of assumptions.
Whatever they left to you, decide, and say what you picked.
Ask a second round, again at most four new questions, only if they ask for more or an answer opens a question you couldn't have asked before.
Never ask the same thing twice.

## An example

Say they shared their café's menu as a PDF, three photos of the room, and a line: "a one-page site for the café, people mostly find us on their phones".
The menu has 40 items and no prices online; the photos are warm and wood-heavy.

Good questions, each from what they sent:

1. "Your menu runs to 40 items. On a phone, what should someone see first?"
   Options: "Today's specials and opening hours (Recommended)", "The full menu, by section", "Where to find you, with a map".
2. "The photos are all warm wood and morning light. Should the page carry that, or stand apart from it?"
   Options: "Carry the room's warmth (Recommended)", "Crisp and quiet, the photos do the talking".
3. "There are no prices online yet. Show them on the site?"
   Options: "Yes, I'll fill them in (Recommended)", "No prices, just the dishes".

Not: "Who is your audience?", which their note already answers, or "What style do you like?", which only free text can answer.
