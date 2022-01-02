package components

// Intentionally verbose as I am still very green when it comes to interacting with the Solana Ecosystem.
// Mostly streaming thoughts as a go along. Hopefully someone else finds this useful.

import (
	"context"
	"eta-multitool/pkg/config"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/sirupsen/logrus"
)

// BuildProgramAccountOptions Build options to filter program accounts by a specific 'Upgrade Authority' public key
// A Solana 'Upgrade Authority' is a key that allows for a program account to be upgraded it is required at Deploy time.
func BuildProgramAccountOptions(pub solana.PublicKey) rpc.GetProgramAccountsOpts {
	offset := uint64(1)
	length := uint64(1024)
	filters := []rpc.RPCFilter{}
	filters = append(filters, rpc.RPCFilter{
		Memcmp: &rpc.RPCFilterMemcmp{
			Offset: 1,
			Bytes:  pub.Bytes(),
		},
	})
	options := rpc.GetProgramAccountsOpts{
		Encoding:   solana.EncodingBase64Zstd,
		Commitment: rpc.CommitmentFinalized,
		DataSlice: &rpc.DataSlice{
			Offset: &offset,
			Length: &length,
		},
		Filters: filters,
	}
	return options
}

// GetAllMetaplexTokenByAuthority returns all accounts owned by the provided program publicKey.
func GetAllMetaplexTokenByAuthority() {
	// Initialize a client with the mainnet-beta endpoint
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	// Log the endpoint
	logrus.WithFields(logrus.Fields{"endpoint": endpoint}).Debug("solana rpc client initialized")

	// 'Base58' instead of standard Base64 encoding?
	// - Don't want 0OIl characters that look the same in some fonts and could be used to create visually identical looking account numbers.
	// - A string with non-alphanumeric characters is not as easily accepted as an account number.
	// - E-mail usually won't line-break if there's no punctuation to break at.
	// - Doubleclicking selects the whole number as one word if it's all alphanumeric.
	//
	// Validate 'Base58' Public keys from string constants.
	// Constants are from Metaplex and GRIM Update Authority

	metaplexProgramKey := solana.MustPublicKeyFromBase58(config.MetaPlexProgramKey)
	grimUpdateAuthorityKey := solana.MustPublicKeyFromBase58(config.GrimUpdateAuthorityKey)

	// 'Program Accounts' are Solana 'Accounts' marked "executable" in its metadata.
	//
	// In the following block we are looking to query against the MetaPlex Candy Machine program account.
	// We are interested in a specific set of downstream accounts so we will need to make use if the additional options.
	// This will allow us just to get back accounts specific to a specific Update Authority Key. ('GRIM')
	programAccountOptions := BuildProgramAccountOptions(grimUpdateAuthorityKey)
	logrus.WithFields(logrus.Fields{"metaplex_program_key": metaplexProgramKey, "update_authority_key": grimUpdateAuthorityKey}).Debug("fetching program accounts")
	resp, err := client.GetProgramAccountsWithOpts(context.Background(), metaplexProgramKey, &programAccountOptions)
	// Handle the error
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("unable to get program accounts")
	}
	logrus.WithFields(logrus.Fields{"count": len(resp)}).Debug("got program accounts")

	// NOTE: This is where the layers of the onion started revealing themselves.
	//
	// I came to understand the linking of the Solana components much better when I walked through them with
	// the https://solscan.io/ tool. I attempted to trace the same process via the UI in the following loop.

	// Now that we have a response containing all accounts owned by the program publicKey we need to handle them. (10k+) at the time of writing.
	for _, value := range resp {
		genesis_signature := GetGenesisSignatureForAddress(client, value.Pubkey)
		token := GetTokenFromTransaction(client, genesis_signature)
		logrus.WithFields(logrus.Fields{
			"token_address": token,
			"public_key":    value.Pubkey},
		).Info(token)
	}
}

// GetTokenFromTransaction
func GetTokenFromTransaction(client *rpc.Client, gen_sig solana.Signature) solana.PublicKey {
	// 'Transactions' are instruction(s) signed by a client using one or more keypairs and executed atomically
	//
	// Get the 'Transaction' from the Genesis (first) 'Signature' passing in empty/default options
	tx, err := client.GetTransaction(context.Background(), gen_sig, &rpc.GetTransactionOpts{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("unable to get signatures for account public key")
	}
	// Parse all 'Account' keys keeping the second entry as that will be our NFT + Metadata Account
	account_keys := tx.Transaction.GetParsedTransaction().Message.AccountKeys
	return account_keys[1]
}

// GetGenesisSignatureForAddress
func GetGenesisSignatureForAddress(client *rpc.Client, pubkey solana.PublicKey) solana.Signature {
	logrus.WithFields(logrus.Fields{
		"public_key": pubkey,
	}).Debug("fetching signatures for address")
	// 'Signatures' are lists of all account public keys referenced by a transaction's instructions.
	// We are interested in this because we want the first Signature
	//
	// Given an individual account lets fetch the 'Signatures' from the Account public key
	sigs, err := client.GetSignaturesForAddress(context.Background(), pubkey)
	// Handle the error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error":      err,
			"public_key": pubkey.String(),
		}).Error("unable to get signatures for account public key")
	}
	// Get the first 'Signature' when the token would have been created
	genesis_signature := sigs[len(sigs)-1].Signature
	logrus.WithFields(logrus.Fields{
		"genesis_signature": genesis_signature,
		"count":             len(sigs),
		"public_key":        pubkey,
	}).Debug("got address signatures")
	return genesis_signature
}
