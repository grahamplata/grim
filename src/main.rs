use clap::{crate_version, App, Arg};
use solana_account_decoder::UiAccountEncoding;
use solana_client::{
    rpc_client::RpcClient,
    rpc_config::{RpcAccountInfoConfig, RpcProgramAccountsConfig},
    rpc_filter::{Memcmp, MemcmpEncodedBytes, RpcFilterType},
};
use solana_transaction_status::UiTransactionEncoding;

// build rpc network configuration
fn build_rpc_cfg(query_key: &str) -> RpcProgramAccountsConfig {
    RpcProgramAccountsConfig {
        account_config: RpcAccountInfoConfig {
            encoding: Some(UiAccountEncoding::Base64Zstd),
            ..RpcAccountInfoConfig::default()
        },
        filters: Some(vec![RpcFilterType::Memcmp(Memcmp {
            offset: 1,
            bytes: MemcmpEncodedBytes::Binary(query_key.to_string()),
            encoding: None,
        })]),
        ..RpcProgramAccountsConfig::default()
    }
}

fn fetch_tokens_by_update_authority() {
    // Setup Communication with a Solana node over RPC.
    let rpc_network = "https://api.mainnet-beta.solana.com";
    let update_authority = "Es1YghGkHZNJ8A9r6oFEHbWsRHbqs4rz6gfkRJ9V4bYf";
    let pubkey = &"metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s"
        .parse()
        .unwrap();
    let client = RpcClient::new(rpc_network.to_owned());
    let cfg = build_rpc_cfg(update_authority);

    println!(
        "Querying Program: {}\nUpdate Authority: {}\nSolana: {}\n",
        pubkey, update_authority, rpc_network,
    );
    // Metaplex Token Metadata Program Public Key
    let metadata_accounts = client
        .get_program_accounts_with_config(pubkey, cfg)
        .expect("could not get program accounts");

    println!("Found {} metadata_accounts\n", metadata_accounts.len());
    for (metadata_address, _account) in metadata_accounts {
        let sigs = client.get_signatures_for_address(&metadata_address);
        if let Err(err) = sigs {
            eprintln!("\ncould not get signatures {} {:?}", pubkey, err);
            continue;
        }

        let sigs = sigs.unwrap();
        if sigs.len() >= 1000 {
            eprintln!("\ntoo many sigs {} {}", pubkey, sigs.len());
            continue;
        }
        if sigs.is_empty() {
            eprintln!("\nnot enough sigs {} {}", pubkey, sigs.len());
            continue;
        }

        let sig = sigs.last().unwrap();
        let sig = sig.signature.parse().unwrap();

        // Returns transaction details for a confirmed transaction
        let tx = client.get_transaction(&sig, UiTransactionEncoding::Base58);
        if let Err(err) = tx {
            eprintln!("\ncouldn't get transaction {} {}", sig, err);
            continue;
        }

        let tx = tx.unwrap().transaction;
        let tx = tx.transaction.decode();
        if tx.is_none() {
            eprintln!("\ncould not decode sig tx {} {}", pubkey, sig);
            continue;
        }

        let tx = tx.unwrap();
        let msg = tx.message();
        if msg.instructions.len() != 5 {
            eprintln!(
                "\ninvalid instruction count {} {}",
                pubkey,
                msg.instructions.len()
            );
            continue;
        }

        let token_address = msg.account_keys.get(1);
        if token_address.is_none() {
            eprintln!("\ncouldn't get token address {}", sig);
            continue;
        }

        let token_address = token_address.unwrap();

        println!("{}", token_address);
    }
}

fn main() {
    let matches = App::new("eta")
        .version(crate_version!())
        .author("Graham Plata <graham.plata@gmail.com>")
        .about("CLI tool to explore the Grim Syndicate and Ethereal Transit Authority ecosystem.")
        .license("MIT")
        .subcommand(
            App::new("fetch").about("fetch token addresses").arg(
                Arg::new("addresses")
                    .short('a')
                    .long("addresses")
                    .about("lists all token addresses"),
            ),
        )
        .subcommand(App::new("watch").about("follow market movement on supported platforms"))
        .subcommand(App::new("floor").about("get the floor price"))
        .subcommand(App::new("community").about("fetch community info"))
        .get_matches();

    // Check for the existence of subcommands
    match matches.subcommand_name() {
        Some("fetch") => fetch_tokens_by_update_authority(),
        Some("community") => println!("fetching community info..."),
        Some("floor") => println!("fetching token floor..."),
        Some("watch") => println!("starting watch..."),
        None => println!("Agent! You forgot to supply a command!"),
        _ => println!("error code 6969, probably nothing"),
    }
}
