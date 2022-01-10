package components

import (
	"context"
	"eta-multitool/pkg/config"
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/near/borsh-go"
)

const EDITION_MARKER_BIT_SIZE uint64 = 248

type Key borsh.Enum

const (
	KeyUninitialized Key = iota
	KeyEditionV1
	KeyMasterEditionV1
	KeyReservationListV1
	KeyMetadataV1
	KeyReservationListV2
	KeyMasterEditionV2
	KeyEditionMarker
)

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

func MetadataDeserialize(data []byte) (Metadata, error) {
	var metadata Metadata
	err := borsh.Deserialize(&metadata, data)
	if err != nil {
		return Metadata{}, fmt.Errorf("failed to deserialize data, err: %v", err)
	}
	// trim null byte
	metadata.Data.Name = strings.TrimRight(metadata.Data.Name, "\x00")
	metadata.Data.Symbol = strings.TrimRight(metadata.Data.Symbol, "\x00")
	metadata.Data.Uri = strings.TrimRight(metadata.Data.Uri, "\x00")
	return metadata, nil
}

type MasterEditionV2 struct {
	Key       Key
	Supply    uint64
	MaxSupply *uint64
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
		panic(err)
	}
	return addr
}

// GetMetadataByAddress
func GetMetadataByAddress(addr string) {
	pubKey := solana.MustPublicKeyFromBase58(addr)
	meta := getMetadata(pubKey)
	endpoint := rpc.MainNetBeta_RPC
	client := rpc.New(endpoint)
	resp, err := client.GetAccountInfo(context.Background(), meta)
	if err != nil {
		println(err)
	}
	spew.Dump(resp.Value.Data.GetBinary())
}
