## Utmost

<p align="center">
  <em>Oswald Chambers' devotional wisdom, right in your terminal.</em>
</p>

## Introduction

Utmost brings the timeless insights of Oswald Chambers' *My Utmost for His Highest* to your command line. It fetches the daily devotional entry from utmost.org and displays it with beautiful terminal styling—perfect for starting your day with reflection before diving into code.

## Demo

```sh
 ⨠ utmost

    My Utmost for His Highest

    │ "Come unto me." — Matthew 11:28

    Are you weary and heavy-laden? Is there a burden weighing on your spirit
    that you cannot shake? Jesus does not say, "Come unto Me and I will give
    you rest for your soul." He says, "Come unto Me," period. The rest comes
    from Him, not from anything you do or feel. The invitation is not to a
    method, but to a Person.

    Bible Reading: Matthew 11:25-30
```

## Installation

### Go

If you have a functional Go environment, you can install with:

```sh
go install github.com/The-Infra-Company/utmost@latest
```

### Source

```sh
git clone git@github.com:The-Infra-Company/utmost.git
cd utmost
make build
```

## Usage

Simply run the command to fetch and display today's devotional:

```sh
utmost
```

The output includes:
- **Title** — The devotional's heading for the day
- **Scripture Verse** — The key verse with reference
- **Body** — Chambers' reflection and insight
- **Bible Reading** — Suggested passage for further study

## Contributing

For bug reports & feature requests, please use the [issue tracker](https://github.com/The-Infra-Company/utmost/issues).

PRs are welcome! We follow the typical "fork-and-pull" Git workflow.
 1. **Fork** the repo on GitHub
 2. **Clone** the project to your own machine
 3. **Commit** changes to your own branch
 4. **Push** your work back up to your fork
 5. Submit a **Pull Request** so that we can review your changes

> [!TIP]
> Be sure to merge the latest changes from "upstream" before making a pull request!
