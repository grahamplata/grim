package components

// Intentionally verbose as I am still very green when it comes to interacting with the Solana Ecosystem.
// Mostly streaming thoughts as a go along. Hopefully someone else finds this useful.

import (
	"context"
	"eta-multitool/pkg/config"
	"eta-multitool/pkg/text"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/sirupsen/logrus"
)

// GetWallet get community wallet balance in SOL
func GetWallet() {
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	pubkey := solana.MustPublicKeyFromBase58(config.GrimCommunityWalletKey)
	balance, err := client.GetBalance(context.Background(), pubkey, rpc.CommitmentFinalized)
	if err != nil {
		fmt.Println("unable to get community wallet balance", err)
	}
	lamports := balance.Value
	sol := text.LamportsToSol(lamports)
	logrus.WithFields(logrus.Fields{
		"sol":      sol,
		"lamports": lamports,
		"message":  fmt.Sprintf("Community wallet balance %d S◎L", sol),
	}).Info(sol)
}
