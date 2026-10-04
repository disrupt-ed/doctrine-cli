# Engineering Decisions

These are defaults, not rules. When a project or the developer says otherwise, they win.

## Understand before you change

Before writing code, read the code around the change. Find out:

- whether the application already does this, fully or partly;
- how similar things are already done here;
- what the language, framework and existing dependencies already provide.

Most unnecessary code comes from skipping this step.

## Choose the simplest solution that works

Go down this list and stop at the first answer that works:

1. Use what the application already has.
2. Use what the language provides.
3. Use what the framework provides.
4. Use a dependency the project already has.
5. Write a small amount of code you own.
6. Add a new dependency.
7. Add new infrastructure (a service, a datastore, a queue).

Each step down adds cost the team pays for as long as the code lives. Moving down needs a reason you can say in one sentence.

## Build for today

Build what the task needs now. Don't add options, hooks, layers or configuration for requirements nobody has asked for.

If you think a future need is likely, mention it in one line. Don't build it.

## Abstract late

Write the direct version first. Add an abstraction (a base class, an interface, a service layer, a generic helper) only when:

- there are already two or three real cases that need it, or
- the direct version is clearly hard to read.

Three similar blocks of code are often better than the wrong abstraction.

## Fit in

Follow the patterns this codebase already uses: naming, file layout, error handling, testing style, how it talks to the database.

If the existing pattern looks wrong, follow it anyway and say what you'd change. Don't introduce a second way of doing the same thing as a side effect of a task.

## Write less code

Less code is less to read, test and maintain. Prefer the solution with fewer lines and fewer files when it is just as clear.

Remove code your change makes dead. Don't leave unused helpers, flags or files behind.

## Test what matters

Test the behavior the task adds or changes, including the edge cases that could actually happen. Follow the project's existing test style and tools.

Don't test the framework. Don't add tests that only restate the implementation.

## Mind security

Treat input from users, URLs, files and other services as untrusted. Watch for injection, unsafe redirects, requests to internal addresses, and leaking private data in responses or logs.

Use the framework's built-in protections rather than your own.

## Mind the data

Think about what happens with real volumes: large tables, many users, slow networks. Mention when a change needs a migration, a backfill or a deploy step.
