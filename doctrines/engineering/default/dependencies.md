# Dependencies

Every dependency is code the team didn't write but has to keep working: upgrades, security fixes, breaking changes, and everything it pulls in.

## Before adding one

Ask, in order:

1. **What is the actual need?** Often it's a small part of what the dependency does.
2. **Does the language or framework already do it?**
3. **Does the project already have a dependency that does it?**
4. **Is it small enough to write and own?** A few dozen clear lines usually are.
5. **Is the dependency healthy?** Check the facts below.

Add it only if the answers point there. Then say why in one sentence.

## Checking health

Look at facts before judging:

- last release and last commit;
- how often it releases;
- number of active maintainers;
- open issues and unanswered critical bugs;
- support for the project's language and framework versions;
- known security advisories;
- how many other dependencies it pulls in.

Report the facts you found, then your judgment. Don't guess. If you couldn't check something, say so.

A dependency with no release in years isn't automatically bad: some libraries are simply finished. It matters when it has open security issues, doesn't support current versions, or nobody answers bugs.

## When the developer names a dependency

If the developer asks for a specific dependency and you see a better option (something already in the project, the framework, or a few lines of code), say so briefly with your reasons before installing it.

If they still want it, install it and move on. Don't argue twice.

## Choosing between candidates

Prefer the one that:

- the ecosystem already treats as standard;
- has fewer transitive dependencies;
- is actively maintained;
- does the job without much more than the job.

Popularity alone isn't a reason. Neither is a long feature list.
