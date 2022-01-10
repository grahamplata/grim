package components

import (
	"context"
	"eta-multitool/pkg/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/near/borsh-go"
	"github.com/sirupsen/logrus"
)

const EDITION_MARKER_BIT_SIZE uint64 = 248

type Key borsh.Enum

type Creator struct {
	Address  solana.PublicKey
	Verified bool
	Share    uint8
}

type Data struct {
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator
}

type Metadata struct {
	Key                 Key
	UpdateAuthority     solana.PublicKey
	Mint                solana.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8
}

// getMetadata
func getMetadata(mint solana.PublicKey) solana.PublicKey {
	TokenMetadataProgramID := solana.MustPublicKeyFromBase58(config.MetaPlexProgramKey)
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			TokenMetadataProgramID.Bytes(),
			mint.Bytes(),
		},
		TokenMetadataProgramID,
	)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("unable to get metadata account")
	}
	return addr
}

func MetadataDeserialize(data []byte) (Metadata, error) {
	var metadata Metadata
	err := borsh.Deserialize(&metadata, data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("failed to deserialize data, err")
	}
	return metadata, nil
}

// GetMetadataByAddress
func GetMetadataByAddress(addr string) {
	mint := solana.MustPublicKeyFromBase58(addr)
	metaDataAccount := getMetadata(mint)
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	resp, err := client.GetAccountInfo(context.Background(), metaDataAccount)
	if err != nil {
		panic(err)
	}
	data := resp.Value.Data.GetBinary()
	metadata, err := MetadataDeserialize(data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("unable to get metadata")
	}
	spew.Dump(metadata)
}
