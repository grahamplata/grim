use gumdrop::Options;
use solana_account_decoder::UiAccountEncoding;
use solana_client::{
    rpc_client::RpcClient,
    rpc_config::{RpcAccountInfoConfig, RpcProgramAccountsConfig},
    rpc_filter::{Memcmp, MemcmpEncodedBytes, RpcFilterType},
};
use solana_transaction_status::UiTransactionEncoding;

// Constants are declared outside all other scopes.
// TODO: Move into external config
// TODO: Maybe extend it to other projects?
// METAPLEX_PUB_KEY is key reference to the Solana PROGRAM ID
const METAPLEX_PUB_KEY: &str = "metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s";
// GRIM_UPDATE_AUTHORITY is key reference to a GRIM Syndicate update authority
const GRIM_UPDATE_AUTHORITY_PUB_KEY: &str = "Es1YghGkHZNJ8A9r6oFEHbWsRHbqs4rz6gfkRJ9V4bYf";

// Define options for the program.
#[derive(Clone, Debug, Options)]
struct AppOptions {
    // Options here can be accepted with any command (or none at all),
    // but they must come before the command name.
    #[options(help = "print help message")]
    help: bool,
    #[options(help = "be verbose")]
    verbose: bool,
    // The `command` option will delegate option parsing to the command type,
    // starting at the first free argument.
    #[options(command)]
    command: Option<Command>,
}

// The set of commands and the options each one accepts.
//
// Each variant of a command enum should be a unary tuple variant with only
// one field. This field must implement `Options` and is used to parse arguments
// that are given after the command name.
#[derive(Clone, Debug, Options)]
enum Command {
    // Command names are generated from variant names.
    // By default, a CamelCase name will be converted into a lowercase,
    // hyphen-separated name; e.g. `FooBar` becomes `foo-bar`.
    //
    // Names can be explicitly specified using `#[options(name = "...")]`
    #[options(help = "fetch 'GRIM' token addresses")]
    Fetch(FetchOpts),
}

// Options accepted for the `fetch` command
#[derive(Clone, Debug, Options)]
struct FetchOpts {
    #[options(help = "print help message")]
    help: bool,
}

// build rpc network configuration
fn build_rpc_cfg(query_key: &str) -> RpcProgramAccountsConfig {
    let cfg = RpcProgramAccountsConfig {
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
    };
    cfg
}

fn fetch_tokens_by_update_authority(
    rpc_network: String,
    program_key: &str,
    update_authority: &str,
) {
    // Setup Communication with a Solana node over RPC.
    let client = RpcClient::new(rpc_network);
    let cfg = build_rpc_cfg(&update_authority);
    let pubkey = &program_key.parse().unwrap();

    // Metaplex Token Metadata Program Public Key
    let metadata_accounts = client
        .get_program_accounts_with_config(pubkey, cfg)
        .expect("could not get program accounts");

    for (metadata_address, _account) in metadata_accounts {
        let sigs = client.get_signatures_for_address(&metadata_address);
        if let Err(err) = sigs {
            eprintln!("\ncould not get signatures {} {:?}", pubkey, err);
            continue;
        }

        // TODO: this will vary per project and should be config
        let sigs = sigs.unwrap();
        if sigs.len() >= 1000 {
            eprintln!("\ntoo many sigs {} {}", pubkey, sigs.len());
            continue;
        }
        if sigs.len() < 1 {
            eprintln!("\nnot enough sigs {} {}", pubkey, sigs.len());
            continue;
        }

        let sig = sigs.last().unwrap();
        let sig = sig.signature.parse().unwrap();

        let tx = client.get_transaction(&sig, UiTransactionEncoding::Base58);
        if let Err(err) = tx {
            eprintln!("\ncouldn't get transaction {} {}", sig, err);
            continue;
        }

        let tx = tx.unwrap().transaction;
        let tx = tx.transaction.decode();
        if let None = tx {
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
        if let None = token_address {
            eprintln!("\ncouldn't get token address {}", sig);
            continue;
        }

        let token_address = token_address.unwrap();

        println!("{}", token_address);
    }
}

fn main() {
    // Parse options from the environment.
    // If there's an error or the user requests help,
    // the process will exit after giving the appropriate response.
    let app_options = AppOptions::parse_args_default_or_exit();
    let default_rpc = "https://api.mainnet-beta.solana.com".to_owned();
    // matchy matchy
    match app_options.clone().command {
        Some(command) => match command {
            Command::Fetch(_app_options) => {
                // TODO: add option to change this
                fetch_tokens_by_update_authority(
                    default_rpc,
                    METAPLEX_PUB_KEY,
                    GRIM_UPDATE_AUTHORITY_PUB_KEY,
                );
            }
        },
        // Default condition
        None => println!("Agent! You forgot to supply a command!"),
    }
}
