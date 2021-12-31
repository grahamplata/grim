package main

//
// Intentionally verbose as I am still very green when it comes to interacting with the Solana Ecosystem.
// Mostly streaming thoughts as a go along. Hopefully someone else finds this useful.
//

import (
	"context"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/sirupsen/logrus"
)

// Executed before the main function and will only be called once. Useful for setting initial application state
func init() {
	if os.Getenv("ENVIRONMENT") == "production" {
		// Log as JSON instead of the default ASCII formatter.
		logrus.SetFormatter(&logrus.JSONFormatter{})
		// Output to stdout instead of the default stderr
		// Can be any io.Writer, see below for File example
		logrus.SetOutput(os.Stdout)
		logrus.WithFields(logrus.Fields{
			"environment": "production",
		}).Info("logger using production config")
	} else {
		// Only log the warning severity or above.
		logrus.SetLevel(logrus.InfoLevel)
		logrus.WithFields(logrus.Fields{
			"environment": "development",
		}).Info("logger using development config")
	}
}

// A Solana 'Upgrade Authority' is a key that allows for a program account to be upgraded it is required at Deploy time.
//
// Build options to filter program accounts by a specific 'Upgrade Authority' public key
func buildProgramAccountOptions(pub solana.PublicKey) rpc.GetProgramAccountsOpts {
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

// main() automatically called when the program is executed
func main() {
	// Initialize a client with the mainnet-beta endpoint
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	// Log the endpoint
	logrus.WithFields(logrus.Fields{"endpoint": endpoint}).Info("solana rpc client initialized")

	// 'Base58' instead of standard Base64 encoding?
	// - Don't want 0OIl characters that look the same in some fonts and could be used to create visually identical looking account numbers.
	// - A string with non-alphanumeric characters is not as easily accepted as an account number.
	// - E-mail usually won't line-break if there's no punctuation to break at.
	// - Doubleclicking selects the whole number as one word if it's all alphanumeric.
	//
	// Validate 'Base58' Public keys from string constants.
	// Constants are from Metaplex and GRIM Update Authority
	metaplexProgramKey := solana.MustPublicKeyFromBase58(MetaPlexProgramKey)
	grimUpdateAuthorityKey := solana.MustPublicKeyFromBase58(GrimUpdateAuthorityKey)

	// 'Program Accounts' are Solana 'Accounts' marked "executable" in its metadata.
	//
	// In the following block we are looking to query against the MetaPlex Candy Machine program account.
	// We are interested in a specific set of downstream accounts so we will need to make use if the additional options.
	// This will allow us just to get back accounts specific to a specific Update Authority Key. ('GRIM')
	programAccountOptions := buildProgramAccountOptions(grimUpdateAuthorityKey)
	logrus.WithFields(logrus.Fields{"metaplex_program_key": metaplexProgramKey, "update_authority_key": grimUpdateAuthorityKey}).Info("fetching program accounts")
	resp, err := client.GetProgramAccountsWithOpts(context.Background(), metaplexProgramKey, &programAccountOptions)
	// Handle the error
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("unable to get program accounts")
	}
	logrus.WithFields(logrus.Fields{"count": len(resp)}).Info("got program accounts")

	// NOTE: This is where the layers of the onion started revealing themselves.
	//
	// I came to understand the linking of the Solana components much better when I walked through them with
	// the https://solscan.io/ tool. I attempted to trace the same process via the UI in the following loop.

	// Now that we have a response containing all accounts owned by the program publicKey we need to handle them. (10k+) at the time of writing.
	for i, value := range resp {
		logrus.WithFields(logrus.Fields{"index": i, "public_key": value.Pubkey}).Info("fetching signatures for address")

		// 'Signatures' are lists of all account public keys referenced by a transaction's instructions.
		// We are interested in this because we want the first Signature
		//
		// Given an individual account lets fetch the 'Signatures' from the Account public key
		sigs, err := client.GetSignaturesForAddress(context.Background(), value.Pubkey)
		// Handle the error
		if err != nil {
			logrus.WithFields(logrus.Fields{"index": i, "error": err, "public_key": value.Pubkey.String()}).Error("unable to get signatures for account public key")
		}
		// Get the first 'Signature' when the token would have been created
		genesis_signature := sigs[len(sigs)-1].Signature
		logrus.WithFields(logrus.Fields{"index": i, "genesis_signature": genesis_signature, "count": len(sigs), "public_key": value.Pubkey}).Info("got address signatures")

		// 'Transactions' are instruction(s) signed by a client using one or more keypairs and executed atomically
		//
		// Get the 'Transaction' from the Genesis (first) 'Signature' passing in empty/default options
		tx, err := client.GetTransaction(context.Background(), genesis_signature, &rpc.GetTransactionOpts{})
		if err != nil {
			logrus.WithFields(logrus.Fields{"index": i, "error": err, "public_key": value.Pubkey.String()}).Error("unable to get signatures for account public key")
		}
		// Parse all 'Account' keys keeping the second entry as that will be our NFT + Metadata Account
		account_keys := tx.Transaction.GetParsedTransaction().Message.AccountKeys
		token_address := account_keys[1]
		logrus.WithFields(logrus.Fields{"index": i, "token_address": token_address, "public_key": value.Pubkey}).Info("got token address")
	}
}
