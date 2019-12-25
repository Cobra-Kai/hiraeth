## Hiraeth

> hiraeth, n: homesickness for a home to which you cannot return, or for a home which may have never been.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Building](#building)
- [Usage](#usage)
- [Support](#support)
- [Contributing](#contributing)

## Introduction

> Cyberspace. A consensual hallucination experienced daily by billions of
> legitimate operators, in every nation, by children being taught mathematical
> concepts... A graphic representation of data abstracted from the banks of
> every computer in the human system. Unthinkable complexity. Lines of light
> ranged in the nonspace of the mind, clusters and constellations of data. Like
> city lights, receding.
> -- <cite>[William Gibson][1]</cite>

Hiraeth is a text-based virtual reality that allows multiple people to engage
in roleplaying, adventuring, and story-telling. The inspiration is taken from
the MUDs, MUCKs, MOOs, MUSHs, and other virtual worlds created in the last four
decades.

## Features

- It's written in Go.
- No other features, only meets basic functional requirements.

## Building

### Build Requirements

go1.13 or later.

### Download source

`go get https://github.com/Cobra-Kai/hiraeth`

### Install

`go install`

### Docker

Optionally, the server may be placed into a container to ease deployment.

    docker build --tag mud .
    docker run mud

NOTE: the above is kind of useless without bindings to see the database and logs.

TODO: need to include a working Dockerfile with the distribution.
TODO: automated testing on every merge [CI](https://en.wikipedia.org/wiki/Continuous_integration)
TODO: distribute **tested** releases on Docker Hub.

## Usage

`hiraeth -http localhost:8080 -telnet :4000`

TODO: provide instructions on how to shut down server cleanly.

### Configuration

TBD

### Back-ups

TBD : describe how to back-up the database.

### Starting for the first time

- Telnet to server, login and create your first account

TODO: provide instructions on how to grant admin access to the first account.

## Support

Please [open an issue](https://github.com/Cobra-Kai/hiraeth/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/).
Create a branch, add commits, and [open a pull request](https://github.com/Cobra-Kai/hiraeth/compare/).

## License

```
Copyright (c) 2019 Jon Mayo <jon@rm-f.net>

Permission to use, copy, modify, and distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
```

## Links & Citations

[1]: Gibson, William (1984). Neuromancer. p. 69
