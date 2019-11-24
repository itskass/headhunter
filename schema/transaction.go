package schema

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type Txdata struct {
	AccountNonce uint64 `json:"nonce"    gencodec:"required"`
	Price        string `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64 `json:"gas"      gencodec:"required"`
	Recipient    string `json:"to"       rlp:"nil"` // nil means contract creation
	Sender       string `json:"from"`
	Amount       string `json:"value"    gencodec:"required"`
	Payload      []byte `json:"input"    gencodec:"required"`
	Size         string `json:"size"`

	// Signature values
	V string `json:"v" gencodec:"required"`
	R string `json:"r" gencodec:"required"`
	S string `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash string `json:"hash" rlp:"-"`
}

func NewTxData(tx *types.Transaction) *Txdata {
	signer := types.NewEIP155Signer(tx.ChainId())
	v, r, s := tx.RawSignatureValues()
	msg, _ := tx.AsMessage(signer)
	from := msg.From()

	return &Txdata{
		Hash:         tx.Hash().String(),
		AccountNonce: tx.Nonce(),
		Price:        hexbig(tx.GasPrice()),
		GasLimit:     tx.Gas(),
		Recipient:    tx.To().String(),
		Sender:       from.String(),
		Amount:       hexbig(tx.Value()),
		Payload:      tx.Data(),
		Size:         tx.Size().String(),
		R:            hexbig(r),
		S:            hexbig(s),
		V:            hexbig(v),
	}
}

func hexbig(b *big.Int) string {
	return fmt.Sprintf("0x%x", b)
}
