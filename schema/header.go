package schema

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Header struct {
	ParentHash  string `json:"parentHash"       gencodec:"required"`
	UncleHash   string `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    string `json:"miner"            gencodec:"required"`
	Root        string `json:"stateRoot"        gencodec:"required"`
	TxHash      string `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash string `json:"receiptsRoot"     gencodec:"required"`
	Difficulty  string `json:"difficulty"       gencodec:"required"`
	Number      uint64 `json:"number"           gencodec:"required"`
	GasLimit    uint64 `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64 `json:"gasUsed"          gencodec:"required"`
	Time        uint64 `json:"timestamp"        gencodec:"required"`
	Extra       []byte `json:"extraData"        gencodec:"required"`
	MixDigest   string `json:"mixHash"`
	Nonce       uint64 `json:"nonce"`
}

func NewHeader(h *types.Header) *Header {
	return &Header{
		ParentHash:  h.ParentHash.String(),
		UncleHash:   h.UncleHash.String(),
		Coinbase:    h.Coinbase.String(),
		Root:        h.Root.String(),
		TxHash:      h.TxHash.String(),
		ReceiptHash: h.ReceiptHash.String(),
		Difficulty:  hexbig(h.Difficulty),
		Number:      h.Number.Uint64(),
		GasLimit:    h.GasLimit,
		GasUsed:     h.GasUsed,
		Time:        h.Time,
		Extra:       h.Extra,
		MixDigest:   h.MixDigest.String(),
		Nonce:       h.Nonce.Uint64(),
	}
}
