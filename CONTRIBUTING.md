# Contributing

## Workflow / Etiquette

The usual pull request work flow applies here with the following specificities:

1. Annotate pull requests with their WIP status. A label or title prefix is fine.
   WIP pull requests are unstable and can be rebased. Non-WIP pull requests are
   append-only and generally safe to merge, posisbly with the author's
   coordination.
1. Give [good feedback][good-feedback] in code reviews (and in general). That
   is, it should be:
   * Actionable
   * Specific
   * Kind (not the same as "positive" -- just coming from a place of empathy)

## Pedantry

These things may seem silly at first, but over time it'll help contributors keep
their sanity as the code base ages. Most of these points are chosen to optimize
for later comprehension or commit archeology instead of for ease-of-committing.
Code is read hundreds of times but only written once.

### Dos

#### Commits

* Write commit messages as per Tim Pope's guidelines [here][commit-messages] ideally
  with detailed explanatory text to commits to provide needed context. Imagine
  `git blame`-ing any of these lines 6 months later; the commit message should
  provide orientation.
* Aim to have every commit pass their tests, as opposed to just the tip of your branch.
  Let's make `git bisect` useful again. See [Source Code History
  Integrity][code-history] and [Git Workflow][git-workflow] for more on this.
* Organize commits into logical, working, independent additions instead of a
  stream of thought. Make generous use of `git rebase -i`, if needed, to
  re-order, merge, and split commits into such components.

#### Code

* Use unix line endings.
* Define all dependencies explicitly, as opposed to relying on dependencies
  having already been loaded elsewhere. For example, in Ruby, never use
  `Bundler.require` -- prefer `Bundler.setup`.
* Wrap all lines as close to 80 chars or less as possible.
* Dependency inject dependencies with sane in-line defaults instead of
  hard-coded constants, singletons, or globals.
* All (at least non-internal) classes have documentation. Classes are commonly
  the first entry point into a code base, often for an on-call engineer
  responding to an exception, so provide enough information to orient
  first-time readers.
* Ensure that the READMEs are still relevant after your change.

### Do nots

* Do not commit failing tests.
* Do not commit pending tests.
* Do not commit commented-out code.
* Do not commit trailing whitespace.
* Do not commit mixed spacing.
* Do not load the full application for unit tests, just the direct component
  itself, which should in turn load it's direct dependencies.
* Do not rely on dependencies included circumstantially (e.g. by a component
  loaded earlier).
* Do not commit artifacts that can be deterministically generated from
  committed source. E.g.: don't commit `.jar` files or compiled assets but do
  commit `Gemfile.lock`.
  * A notable exception here is if there isn't an opportunity to generate these
    files before typical use (such as, as a dependency).

[good-feedback]: http://www.pechakucha.org/presentations/the-most-valuable-skill
[commit-messages]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html
[code-history]: https://www.destroyallsoftware.com/screencasts/catalog/source-code-history-integrity
[git-workflow]: https://www.destroyallsoftware.com/screencasts/catalog/git-workflow
