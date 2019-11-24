package schema

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Block struct {
	Hash             string
	Header           *Header
	Transactions     []*Txdata
	TransactionCount int
}

func NewBlock(b *types.Block) *Block {
	tx := b.Transactions()
	return &Block{
		Hash:             b.Hash().String(),
		Header:           NewHeader(b.Header()),
		Transactions:     parseTransactions(b.Transactions()),
		TransactionCount: len(tx),
	}
}

func parseTransactions(transactions []*types.Transaction) []*Txdata {
	var txs []*Txdata
	for _, tx := range transactions {
		txs = append(txs, NewTxData(tx))
	}
	return txs
}
