# ETA Multitool - A Grim Syndicate CLI

---

- [ETA Multitool - A Grim Syndicate CLI](#eta-multitool---a-grim-syndicate-cli)
- [What’s The Grim Syndicate?](#whats-the-grim-syndicate)
  - [Links](#links)
- [ETA Multi-Tool](#eta-multi-tool)
  - [Project Setup](#project-setup)
  - [Usage](#usage)
  - [Configuration](#configuration)
  - [Commands](#commands)
    - [Fetch](#fetch)
  - [Flags](#flags)
    - [Logging](#logging)
      - [None](#none)
      - [Text](#text)
      - [JSON](#json)
  - [Project Tree](#project-tree)
  - [Notes](#notes)
    - [Why is there a Rust Folder?](#why-is-there-a-rust-folder)

---

# What’s The Grim Syndicate?

<p align="center">
  <img style="max-width:150px;" src="./docs/318.jpg" />
</p>
<p align="center">
  10,000 generative NFTs on the #Solana blockchain.
</p>

> The Grim Syndicate are an elite collective responsible for ferrying fickle Souls from across the dimensional spectrum (even those hard-to-reach pocket universes) to arrive at their final destination. At the Ethereal Transit Authority death doesn't have to mean a dead-end. -- https://grimsyndicate.com/

PS... Don't forget your [ID badge!](https://grimsyndicate.id/)

## Links

- [Website](https://grimsyndicate.com/)
- [Twitter](https://twitter.com/Grim__Syndicate)
- [Discord](https://discord.gg/xeHPSUhUv7)

# ETA Multi-Tool

> Part of the standard issue Ethereal Transit Authority Toolbelt

A simple CLI tool to explore the [Grim Syndicate](https://grimsyndicate.com/) and Ethereal Transit Authority ecosystem.

## Project Setup

For most of the project I have intentionally heavily commented as I am still very green when it comes to interacting with the Solana Ecosystem. I tried to keep it as a set of streaming thoughts as I went along. Hopefully someone else finds this useful.

```shell
git clone git@github.com:grahamplata/grim.git
cd grim
go build -o eta-multitool -v .
```

## Usage

```shell
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  eta-multitool [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  fetch       A brief description of your command
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.eta.yaml)
  -h, --help            help for eta-multitool
  -o, --output string   output type of the cli (default is ...)

Use "eta-multitool [command] --help" for more information about a command.
```

## Configuration

```bash
# coming soon to a $HOME/dir near you
```

## Commands

### Fetch

## Flags

### Logging

> Support a wide variety of logging formats for every usecase.

#### None

> Output as a minimal standard out

```bash
# none - eta-multitool fetch
7k1kyD37tLv528fWSs3wujtHFemwMGmTTxX5dqp8fN1A
```

#### Text

> Output as a detailed text filed

```bash
# text - eta-multitool fetch --output="text"
INFO[2022-01-01T13:07:43-05:00] logger using text config
INFO[2022-01-01T13:08:52-05:00] 6SjSBmT2cd1sgS6yqiHVhDroWBE11f72PopyPgkb5AMQ  index=0 public_key=DWviW6d8AK4ksTL9wJvXy8XBqevtGrY5UQgmTRdiQHqd token_address=6SjSBmT2cd1sgS6yqiHVhDroWBE11f72PopyPgkb5AMQ
```

#### JSON

> Output as a json object

```bash
# json - eta-multitool fetch --output="json"
{"level":"info","msg":"logger using json config","time":"2022-01-01T13:05:46-05:00"}
{"index":0,"level":"info","msg":"DU6TVCLMYwoomiG2EvY8ECCFvquc5iefwdJbDMYtjr7T","public_key":"3X7HzqxBquHf8Sgqd5Tzi93b3frP74k68JeHAhVTrnNb","time":"2022-01-01T13:06:45-05:00","token_address":"DU6TVCLMYwoomiG2EvY8ECCFvquc5iefwdJbDMYtjr7T"}
```

## Project Tree

> Project structure

```shell
# tree -L 3
.
├── LICENSE
├── README.md
├── cmd
│   ├── fetch.go
│   └── root.go
├── deprecated
│   └── rust
│       ├── Cargo.lock
│       ├── Cargo.toml
│       ├── README.md
│       ├── src
│       └── target
├── docs
│   └── grim-318.jpg
├── go.mod
├── go.sum
├── main.go
└── pkg
    ├── components
    │   └── token.go
    ├── config
    │   ├── constants.go
    │   └── logging.go
    └── text
```

## Notes

### Why is there a Rust Folder?

> I originally started with rust to tinker. If you are still interested the code lives here: [Link](./deprecated/rust/README.md)
