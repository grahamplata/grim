# Grim Syndicate CLI

A simple CLI tool to explore the [Grim Syndicate](https://grimsyndicate.com/) and Ethereal Transit Authority ecosystem.

---

- [Grim Syndicate CLI](#grim-syndicate-cli)
  - [What’s The Grim Syndicate?](#whats-the-grim-syndicate)
    - [Links](#links)
  - [Prerequisites](#prerequisites)
  - [Install](#install)
  - [Usage](#usage)
    - [Help](#help)
    - [Fetch](#fetch)
  - [Personal Goals](#personal-goals)

---

## What’s The Grim Syndicate?

10,000 generative NFTs on the #Solana blockchain.

> The Grim Syndicate are an elite collective responsible for ferrying fickle Souls from across the dimensional spectrum (even those hard-to-reach pocket universes) to arrive at their final destination. At the Ethereal Transit Authority death doesn't have to mean a dead-end. -- https://grimsyndicate.com/

PS... Don't forget your [ID badge!](https://grimsyndicate.id/)

![](./docs/grim-318.jpg)

### Links

- [Website](https://grimsyndicate.com/)
- [Twitter](https://twitter.com/Grim__Syndicate)
- [Discord](https://discord.gg/xeHPSUhUv7)

## Prerequisites

```bash
# Install Rust
# - https://www.rust-lang.org/tools/install
# - https://forge.rust-lang.org/infra/other-installation-methods.html

curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

## Install

```bash
git clone https://github.com/grahamplata/sterling-pigeon.git
cd sterling-pigeon/grims
cargo run
# building
cargo build
cargo install --path .
grim
```

> To be able to run `grim`, you must be connected to the internet; you can read all content offline, however!

## Usage

Use this tool to query the **Ethereal Transit Authority** **(ETA)**.

### Help

```bash
grim --help
Usage: grim [OPTIONS]

Optional arguments:
  -h, --help     print help message
  -v, --verbose  be verbose

Available commands:
  fetch  fetch 'GRIM' tokens
```

### Fetch

```bash
grim fetch
Fetching 'GRIM' on https://api.mainnet-beta.solana.com
9vDQ3yRoZJAiy2nrksWawnQzHB2TRd8Tvzp9A4VFWaaZ
C1TVanivQMcwnpbfL34VqNmrDd7kZSGai8Z3LVvfFxgN
7bTLUnhkaRWzF5y2pDf9JMwMXHxBL2jg5M6PAexqULYZ
```

## Personal Goals

Learn and contribute to the 'GRIM' community.

- Community building
- Blockchain
- Solana
- Rust (first project)
- NFTs
