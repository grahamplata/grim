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

# Solana on linux
# https://github.com/solana-labs/solana#1-install-rustc-cargo-and-rustfmt
sudo apt-get update
sudo apt-get install libssl-dev libudev-dev pkg-config zlib1g-dev llvm clang make
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
Usage: target/debug/grim [OPTIONS]

Optional arguments:
  -h, --help                 print help message
  -v, --verbose              be verbose
  -r, --rpc-url RPC          rpc network (default: https://api.mainnet-beta.solana.com)
  -p, --program-id metaplex  program id (default: metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s)

Available commands:
  fetch  fetch token addresses
```

### Fetch

```bash
grims fetch --help
Usage: target/debug/grim fetch [OPTIONS]

Optional arguments:
  -h, --help  print help message
  -u, --update-authority grims
              update authority address (default: Es1YghGkHZNJ8A9r6oFEHbWsRHbqs4rz6gfkRJ9V4bYf)
```

```bash
grims fetch
Starting fetch on metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s using Es1YghGkHZNJ8A9r6oFEHbWsRHbqs4rz6gfkRJ9V4bYf
Found 10000 metadata_accounts

EZJjAPxKhe4c7Xs1kfLBmwKqnY9owmXihVg2DciiMjrY
8FdMo6dk8CFj1cYN2FoAXEZjHhNP1jcXhn9QbwUEbtaH
B2CAP934Qdive3WQtzMDfYP3YNhmQHQ8sVJEeBrrUi4j
6iAM7CH2KgsnNLHhpEj1VghBRfK5N3v8Ki4SmWN6P4Da
HM5F5Vm28jvfUzpQTj5kMJPPdHrWYQZryPDutrhqmR76
```

## Personal Goals

Learn and contribute to the 'GRIM' community.

- Community building
- Blockchain
- Solana
- Rust (first project)
- NFTs
