package gather

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/globalsign/mgo/bson"

	"github.com/itskass/headhunter/schema"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/globalsign/mgo"
)

type Options struct {
	// Target of this gather, target can be a hash
	// block number or "latest".
	// REQUIRED
	Target bson.M
	// Client is used to communicate with the connected
	// nodes rpc.
	// REQUIRED
	Client *ethclient.Client
	// DB is used to store gathered blocks
	// REQUIRED
	DB *mgo.Database
	// ShouldSynchronize if true will download missing blocks
	// using provided sync options.
	ShouldSynchronize bool
	// SyncOptions contain head and tail values to synchronize
	// between
	SyncOptions SyncOptions
	// GetAncestors if true, the target blocks ancestors will be
	// gathered recursively.
	GetAncestors bool
}

func Blocks(opts *Options) {

	var (
		blockData *types.Block
		err       error
		target    = opts.Target
		c         = opts.Client
		db        = opts.DB
	)

	// request the block with this hash via rpc
	if _hash, ok := target["hash"]; ok {
		hash := _hash.(string)
		Log.Println("rpc: requesting block by hash", hash)
		if blockData, err = blockByHash(hash, c); err != nil {
			log.Println("rpc: err:", err)
			return
		}
	}

	// or request by number via rpc
	if _number, ok := target["number"]; ok {
		number := _number.(uint64)
		Log.Println("rpc: request block by number", number)
		if blockData, err = blockByNumber(number, c); err != nil {
			log.Println("rpc: err", err)
			return
		}
	}

	if _, ok := target["latest"]; ok {
		Log.Println("rpc: request latest block")
		blockData, err = c.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Println("rpc: err", err)
			return
		}
	}

	// convert block to mongodb friendly schema block
	// which uses hex values for hashes, addresses and
	// big.Ints
	block := schema.NewBlock(blockData)

	// save block to database
	_, err = db.C("blocks").Upsert(bson.M{"hash": block.Hash}, block)
	if err != nil {
		log.Println("db: err:", err)
		return
	}
	Log.Println("db: saved block", block.Hash)

	// get ancestors by getting ancestors if
	// ancestor exists
	if opts.GetAncestors {
		if !hasAncestor(block, db) {
			Log.Println("sync: complete: no more missing ancestors")
			return
		}
		Log.Println("sync: unsynced ancestor found, requesting...")
		opts.Target = bson.M{"hash": block.Header.ParentHash}
		Blocks(opts)
	}

	// synchronize
	if opts.ShouldSynchronize {
		// disable only need to synchronize once
		opts.ShouldSynchronize = false
		// default sync options then synchronize from
		// current block to root
		if opts.SyncOptions.Head == 0 {
			opts.SyncOptions.Head = block.Header.Number
			opts.SyncOptions.Tail = 0
		}
		// Synchronize will download blocks between
		// the provided tail and head.
		Synchronize(opts)
	}
}

func hasAncestor(b *schema.Block, db *mgo.Database) bool {
	// check if block is genesis
	if b.Header.Number < 1 {
		return false
	}

	// check if blocks parent is already owned
	return !owned(bson.M{"hash": b.Header.ParentHash}, db)
}

func blockByHash(hash string, c *ethclient.Client) (*types.Block, error) {
	_hash := common.HexToHash(hash)
	return c.BlockByHash(
		context.Background(),
		_hash,
	)
}

func blockByNumber(number uint64, c *ethclient.Client) (*types.Block, error) {
	_number := new(big.Int).SetUint64(number)
	return c.BlockByNumber(
		context.Background(),
		_number,
	)
}

func owned(query bson.M, db *mgo.Database) bool {
	n, _ := db.C("blocks").Find(query).Count()
	return n > 0
}
