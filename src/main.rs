mod state;

use std::{io::Read, str::FromStr};

use anchor_lang::AnchorDeserialize;
use clap::{crate_authors, crate_version, App, ArgMatches};
use solana_account_decoder::UiAccountEncoding;
use solana_client::{
    rpc_client::RpcClient,
    rpc_config::{RpcAccountInfoConfig, RpcProgramAccountsConfig},
    rpc_filter::{Memcmp, MemcmpEncodedBytes, RpcFilterType},
};
use solana_program::borsh::try_from_slice_unchecked;
use solana_sdk::pubkey::Pubkey;
use solana_sdk::{account::ReadableAccount, native_token::lamports_to_sol};
use spl_token_metadata::state::Metadata;

// Setup Communication with a Solana node over RPC.
fn build_rpc_client() -> RpcClient {
    let rpc_network = state::RPC_NETWORK;
    RpcClient::new(rpc_network.to_string())
}

// Setup configuration Rpc Program Accounts Config
fn build_rpc_client_cfg() -> RpcProgramAccountsConfig {
    RpcProgramAccountsConfig {
        account_config: RpcAccountInfoConfig {
            encoding: Some(UiAccountEncoding::Base64Zstd),
            ..RpcAccountInfoConfig::default()
        },
        filters: Some(vec![RpcFilterType::Memcmp(Memcmp {
            offset: 1,
            bytes: MemcmpEncodedBytes::Binary(state::PROGRAM_UPGRADE_AUTHORITY.to_string()),
            encoding: None,
        })]),
        ..RpcProgramAccountsConfig::default()
    }
}

fn fetch_tokens_by_update_authority() {
    let pubkey = &state::PROGRAM_PUBLIC_KEY
        .parse::<Pubkey>()
        .expect("Failed to parse Metaplex Program Id");
    let client = build_rpc_client();
    let config = build_rpc_client_cfg();

    // Metaplex Token Metadata Program Public Key
    let metadata_accounts = client
        .get_program_accounts_with_config(pubkey, config)
        .expect("unable to fetch program accounts");

    println!("Found {} metadata_accounts\n", metadata_accounts.len());
    for (_, account) in metadata_accounts {
        let metadata: Metadata = try_from_slice_unchecked(&account.data)
            .expect("unable to get mint address from account");
        println!("{}", metadata.mint.to_string());
    }
}

fn fetch_arweave_metadata(uri: &str) -> Result<state::Grim, Box<dyn std::error::Error>> {
    let mut res = reqwest::blocking::get(uri)?;
    let mut body = String::new();
    res.read_to_string(&mut body)?;
    let deserialized: state::Grim = serde_json::from_str(&body).unwrap();
    Ok(deserialized)
}

fn fetch_token_metadata(mint: &str) {
    let pubkey = Pubkey::from_str(&mint).expect("msg");
    let client = build_rpc_client();
    let account = client
        .get_account(&pubkey)
        .expect("could not fetch metadata account");

    let mut buf = account.data();
    let metadata = Metadata::deserialize(&mut buf).expect("could not deserialize metadata");
    println!("Name: {}", metadata.data.name);
    println!("Symbol: {}", metadata.data.symbol);
    println!("Metadata: {}", metadata.data.uri);
    println!("Update Authority: {}", metadata.update_authority);
    println!("Key: {:?}", metadata.key);
    println!("Mint: {}", metadata.mint);
    println!("Primary Sale Happened: {}", metadata.primary_sale_happened);
    println!("Mutable: {}", metadata.is_mutable);
    let grim = fetch_arweave_metadata(&metadata.data.uri).expect("unable to fetch metadata");
    println!("{:#?}", grim);
}

// Community
fn fetch_community_wallet_balance() {
    let pubkey = &state::COMMUNITY_WALLET_PUBLIC_KEY.parse().unwrap();
    let client = build_rpc_client();
    let account_balance = client
        .get_balance(pubkey)
        .expect("could not get account balance");
    println!("'GRIM' S◎L Balance: {}", lamports_to_sol(account_balance));
}

// Command responses
fn handle_command_none() {
    println!("Agent! You forgot to supply a command!")
}

fn handle_unreachable_command() {
    println!("error code 6969, probably nothing")
}

// Command matching
fn handle_command_matches(matches: ArgMatches) {
    match matches.subcommand() {
        Some(("fetch", fetch_matches)) => match fetch_matches.subcommand() {
            Some(("metadata", val)) => fetch_token_metadata(val.value_of("metadata").unwrap()),
            Some(("tokens", _)) => fetch_tokens_by_update_authority(),
            None => handle_command_none(),
            _ => handle_unreachable_command(),
        },
        Some(("community", community_matches)) => match community_matches.subcommand() {
            Some(("wallet", _)) => fetch_community_wallet_balance(),
            Some(("holders", _)) => println!("fetching community holders..."),
            None => handle_command_none(),
            _ => handle_unreachable_command(),
        },
        Some(("floor", _floor_matches)) => println!("fetching token floor..."),
        Some(("watch", _watch_matches)) => println!("starting watch..."),
        None => handle_command_none(),
        _ => handle_unreachable_command(),
    }
}

// main
fn main() {
    let matches = App::new("eta")
        .version(crate_version!())
        .author(crate_authors!())
        .about("CLI tool to explore the Grim Syndicate and Ethereal Transit Authority ecosystem.")
        .license("MIT")
        .subcommand(
            App::new("fetch")
                .about("fetch all token addresses")
                .subcommand(App::new("metadata").about("fetch metadata"))
                .subcommand(App::new("tokens").about("fetches all token addresses")),
        )
        .subcommand(
            App::new("community")
                .about("fetch community info")
                .subcommand(App::new("wallet").about("fetch community wallet balance"))
                .subcommand(App::new("holders").about("fetch token holders")),
        )
        .subcommand(App::new("watch").about("follow market movement on supported platforms"))
        .subcommand(App::new("floor").about("get the floor price"))
        .get_matches();
    handle_command_matches(matches);
}
