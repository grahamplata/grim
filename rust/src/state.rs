use serde::{Deserialize, Serialize};

// Solana Configuration
// https://docs.solana.com/clusters
// Solana RPC service running on nyc4.rpcpool.com in region US on network mainnet.
pub const RPC_NETWORK: &str = "https://api.mainnet-beta.solana.com";

// 'GRIM' community wallet public key
pub const COMMUNITY_WALLET_PUBLIC_KEY: &str = "RTp26f9wY2fXxeWRE7FkS9iVrsuxgdUJfDYH8GgoBH9";

// Developers can deploy on-chain programs (often called smart contracts elsewhere) with the Solana tools.
// https://docs.solana.com/cli/deploy-a-program
//
// Token Metadata Program Account public key
pub const PROGRAM_PUBLIC_KEY: &str = "metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s";
// 'GRIM' Program Upgrade Authority
pub const PROGRAM_UPGRADE_AUTHORITY: &str = "Es1YghGkHZNJ8A9r6oFEHbWsRHbqs4rz6gfkRJ9V4bYf";

#[derive(Serialize, Deserialize, Debug)]
pub struct Grim {
    pub name: String,
    pub symbol: String,
    pub description: String,
    pub seller_fee_basis_points: i16,
    pub image: String,
    pub external_url: String,
    pub collection: Collection,
    pub attributes: Vec<Attributes>,
    pub properties: Properties,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Collection {
    pub name: String,
    pub family: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Attributes {
    pub trait_type: String,
    pub value: Trait,
}

#[derive(Serialize, Deserialize, Debug)]
#[serde(untagged)]
pub enum Trait {
    Numeric(i32),
    Textual(String),
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Properties {
    pub files: Vec<File>,
    pub creators: Vec<Creator>,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct File {
    pub uri: String,
    pub r#type: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Creator {
    pub address: String,
    pub share: u8,
}
