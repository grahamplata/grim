package main

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

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

// main
func main() {
	// Initialize a websocket client with the mainnet-beta endpoint
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	metaplexProgramKey := solana.MustPublicKeyFromBase58(MetaPlexProgramKey)
	grimUpdateAuthorityKey := solana.MustPublicKeyFromBase58(GrimUpdateAuthorityKey)
	programAccountOptions := buildProgramAccountOptions(grimUpdateAuthorityKey)

	// Make GetProgramAccounts call
	resp, err := client.GetProgramAccountsWithOpts(context.Background(), metaplexProgramKey, &programAccountOptions)
	if err != nil {
		panic(err)
	}

	// Layers of the onion
	for _, value := range resp {
		// spew.Dump("value:", value)
		sigs, err := client.GetSignaturesForAddress(context.Background(), value.Pubkey)
		if err != nil {
			panic(err)
		}
		// spew.Dump("sigs:", sigs)
		genesis_signature := sigs[len(sigs)-1].Signature
		tx, err := client.GetTransaction(context.Background(), genesis_signature, &rpc.GetTransactionOpts{})
		if err != nil {
			panic(err)
		}
		token_address := tx.Transaction.GetParsedTransaction().Message.AccountKeys
		fmt.Println(token_address[1])
	}
}
