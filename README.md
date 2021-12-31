# ETA Multitool - A Grim Syndicate CLI

[![Continuous integration](https://github.com/grahamplata/grim/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/grahamplata/grim/actions/workflows/ci.yml)

---

- [ETA Multitool - A Grim Syndicate CLI](#eta-multitool---a-grim-syndicate-cli)
  - [What’s The Grim Syndicate?](#whats-the-grim-syndicate)
    - [Links](#links)
  - [ETA Tool](#eta-tool)
    - [Setup](#setup)
    - [Usage](#usage)
    - [Configuration](#configuration)
    - [Tree](#tree)
  - [Notes](#notes)
    - [Why is there a Rust Folder?](#why-is-there-a-rust-folder)

---

## What’s The Grim Syndicate?

10,000 generative NFTs on the #Solana blockchain.

> The Grim Syndicate are an elite collective responsible for ferrying fickle Souls from across the dimensional spectrum (even those hard-to-reach pocket universes) to arrive at their final destination. At the Ethereal Transit Authority death doesn't have to mean a dead-end. -- https://grimsyndicate.com/

PS... Don't forget your [ID badge!](https://grimsyndicate.id/)

<p align="center">
  <img src="./docs/grim-318.jpg" />
</p>

### Links

- [Website](https://grimsyndicate.com/)
- [Twitter](https://twitter.com/Grim__Syndicate)
- [Discord](https://discord.gg/xeHPSUhUv7)

## ETA Tool

> Standard issue Ethereal Transit Authority Toolbelt

A simple CLI tool to explore the [Grim Syndicate](https://grimsyndicate.com/) and Ethereal Transit Authority ecosystem.

### Setup

Build it yourself

```shell
git clone git@github.com:grahamplata/grim.git
cd grim
go build -o eta -v .
```

### Usage

```bash
# Sample Output
go run main.go constants.go
INFO[0000] logger using development config               environment=development
INFO[0000] solana rpc client initialized                 endpoint="https://api.mainnet-beta.solana.com"
INFO[0000] fetching program accounts                     metaplex_program_key=metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s update_authority_key=Es1YghGkHZNJ8A9r6oFEHbWsRHbqs4rz6gfkRJ9V4bYf
INFO[0051] got program accounts                          count=19405
INFO[0051] fetching signatures for address               index=0 public_key=AivbYVPoPRX7WuEkba5cjonJQKzno1xpizPbNpSafB7n
INFO[0051] got address signatures                        count=10 genesis_signature=4nK456176oH1RaaVqfz7s4QmcwFKgphB45uWkdHjX7v8JUijrQ7cZXejh13nzJ8MXZx2AYAoUsT9z57HFz1oMMzM index=0 public_key=AivbYVPoPRX7WuEkba5cjonJQKzno1xpizPbNpSafB7n
INFO[0051] got token address                             index=0 public_key=AivbYVPoPRX7WuEkba5cjonJQKzno1xpizPbNpSafB7n token_address=Dnsu6Doj86Yng64ZQrCNRQ2kHZoXFAs5qnKtRQXSgNPP

```

### Configuration

```bash

```

### Tree

```bash

```

## Notes

### Why is there a Rust Folder?

> I originally started with rust to tinker. If you are still interested the code lives here: [Link](./deprecated/rust/README.md)
