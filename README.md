# ETA Multitool - A Grim Syndicate CLI

[![Continuous integration](https://github.com/grahamplata/grim/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/grahamplata/grim/actions/workflows/ci.yml)

---

- [ETA Multitool - A Grim Syndicate CLI](#eta-multitool---a-grim-syndicate-cli)
- [What’s The Grim Syndicate?](#whats-the-grim-syndicate)
  - [Links](#links)
- [ETA Tool](#eta-tool)
  - [Project Setup](#project-setup)
  - [Usage](#usage)
  - [Configuration](#configuration)
  - [Tree](#tree)
  - [Logging](#logging)
    - [None](#none)
    - [Text](#text)
    - [JSON](#json)
  - [Notes](#notes)
    - [Why is there a Rust Folder?](#why-is-there-a-rust-folder)

---

# What’s The Grim Syndicate?

10,000 generative NFTs on the #Solana blockchain.

> The Grim Syndicate are an elite collective responsible for ferrying fickle Souls from across the dimensional spectrum (even those hard-to-reach pocket universes) to arrive at their final destination. At the Ethereal Transit Authority death doesn't have to mean a dead-end. -- https://grimsyndicate.com/

PS... Don't forget your [ID badge!](https://grimsyndicate.id/)

<p align="center">
  <img src="./docs/grim-318.jpg" />
</p>

## Links

- [Website](https://grimsyndicate.com/)
- [Twitter](https://twitter.com/Grim__Syndicate)
- [Discord](https://discord.gg/xeHPSUhUv7)

# ETA Tool

> Standard issue Ethereal Transit Authority Toolbelt

A simple CLI tool to explore the [Grim Syndicate](https://grimsyndicate.com/) and Ethereal Transit Authority ecosystem.

## Project Setup

```shell
git clone git@github.com:grahamplata/grim.git
cd grim
go build -o eta-multitool -v .
```

## Usage

```bash
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

```

## Tree

```bash

```

## Logging

> Support a wide variety of logging formats for every usecase.

### None

Output as a minimal standard out

```bash
# none - eta-multitool fetch
7k1kyD37tLv528fWSs3wujtHFemwMGmTTxX5dqp8fN1A
```

### Text

Output as a detailed text filed

```bash
# text - eta-multitool fetch --output="text"
INFO[2022-01-01T13:07:43-05:00] logger using text config
INFO[2022-01-01T13:08:52-05:00] 6SjSBmT2cd1sgS6yqiHVhDroWBE11f72PopyPgkb5AMQ  index=0 public_key=DWviW6d8AK4ksTL9wJvXy8XBqevtGrY5UQgmTRdiQHqd token_address=6SjSBmT2cd1sgS6yqiHVhDroWBE11f72PopyPgkb5AMQ
```

### JSON

Output as a json object

```bash
# json - eta-multitool fetch --output="json"
{"level":"info","msg":"logger using json config","time":"2022-01-01T13:05:46-05:00"}
{"index":0,"level":"info","msg":"DU6TVCLMYwoomiG2EvY8ECCFvquc5iefwdJbDMYtjr7T","public_key":"3X7HzqxBquHf8Sgqd5Tzi93b3frP74k68JeHAhVTrnNb","time":"2022-01-01T13:06:45-05:00","token_address":"DU6TVCLMYwoomiG2EvY8ECCFvquc5iefwdJbDMYtjr7T"}
```

## Notes

### Why is there a Rust Folder?

> I originally started with rust to tinker. If you are still interested the code lives here: [Link](./deprecated/rust/README.md)
