# Publishing a design as an answer

Read this before you upload a design, replace one, or tell someone how to open it.
One design is one answer, at one link, revised in place round after round.

## Choose the workspace

Call `list_workspaces`.
If there's only one, use it; if the person named one, use that; otherwise ask them once which to use.
You need to be an owner or a contributor there to publish; a viewer only reads.

## Who can open it

A new answer takes its workspace's default, which `get_workspace` gives as `new_answers`, and every reply gives as `who_can_open`.
Tell the person which it is when you hand back the link.

- Never make an answer open to anyone with the link unless the person asks; `set_who_can_open` changes it, and changes the link too.
- To show the design to someone else in a restricted workspace, an owner can add them as a viewer with `share_answer`, which lets them open every answer in that workspace, and only works for someone already invited to ILTTU.
  ILTTU sends no mail, so send them the link yourself, and tell them to sign in with that address.
- If you're not an owner, ask one to add them.

## Pack the page

- A page that is one file goes up as `index.html` or any `.html` file, which becomes the answer's `index.html`.
- A page with other files goes up as a `.zip`, `.tar`, `.tar.gz` or `.tgz`, with `index.html` at its top.
- Link files by relative paths, such as `styles.css` or `option-b.html`, never by absolute paths starting with `/`.

Before you upload, check the page against what ILTTU refuses, so a round isn't lost to it:

- an archive with no `index.html` at its top, unless everything sits in one folder that has one;
- a file with no extension, or one that isn't a web file type;
- a symbolic or hard link, or a path that climbs out of the archive;
- more files, or more bytes once unpacked, than the instance allows, which the refusal names.

And look at it in a local browser where there is one, at desktop and phone width.

## Publish the first round

1. Call `start_upload` with the file's name, its size in bytes and the workspace's ID.
2. Within 5 minutes, send one multipart POST to the `url` it returns, with each of its `fields` in order and then the file as the field named `file`.
3. Read the answer with `get_answer`, using the `id` from the reply, until its `status` is no longer `pending`.

A shell with `curl` and `jq` can send the POST, with the reply saved as `reply.json`:

```bash
args=(); while IFS= read -r kv; do args+=(--form-string "$kv"); done < <(jq -r '.fields | to_entries[] | "\(.key)=\(.value)"' reply.json)
curl -sS --fail-with-body "${args[@]}" -F "file=@design.zip" "$(jq -r .url reply.json)"
```

`--form-string` sends each field as it is, even one that starts with `@` or `<`.
The reply holds a short-lived upload credential: keep it out of your messages to the person, and delete the file once it's sent.

When `get_answer` says `ready`, it gives the answer's `link` and its `version`, which the next round needs.
`failed` gives the reason: fix what it names and start again with `start_upload`, which makes a new answer.

## Options as pages

A round of options is one answer:

- `index.html` is an overview that names each option and links to its page;
- each option is a page of its own beside it, such as `harbour.html` and `market.html`, each linking back to the overview.

Each option's case and its tradeoff go in your reply, not on its page.
Keep each option's name and file name the same in every later round.

When the person picks one, replace the answer with that option alone.
First tell them which options will be gone from the link, and that they're gone for good.
If they want to keep one, publish it as a separate answer before you replace.

## Revise it, round by round

Each round replaces the answer at the same link, keeping its ID and who can open it.

1. Call `replace_answer` with the answer's ID, the `version` from the last `ready` reply, the new file's name and size, and a new token for this round, such as `ferry-round-3-k2p7`.
2. Within 5 minutes, send the multipart POST to the `url` it returns, as for the first round.
3. Call `replace_answer` again with the same token and file, which waits up to 20 seconds and starts nothing new.
   - `ready` gives the new `version`, which the next round sends.
   - `pending` means the link still opens the last round: call it again with the same token.
   - `failed` leaves the last round in place, with the reason: fix it, and start over with a new token.

What the answer served before is gone for good once a replacement completes.

After a round is ready, tell the person the link may show the old page for a moment.

## When a round is refused

On any refusal from `replace_answer`, read the answer with `get_answer`, and go by what it says, never by the refusal's words, which can change.

- If it's `ready` at the version you sent, wait a minute, then try again with a new token.
- If it serves another version, it has changed since you last read it: tell the person, and don't replace it again until they say what to do.
- If it's `failed`, `blocked` or gone, tell the person what `get_answer` says.

## Hand back the link

Each round's reply to the person holds:

- the link, and whether it's ready;
- who can open it;
- for a round of options, each option's case and its main tradeoff;
- your assumptions, in one line;
- what you'd like them to tell you, such as which option to keep or what to change.
