package subscribe

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/globalsign/mgo/bson"
	"github.com/itskass/headhunter/gather"
)

var headerChan = make(chan *types.Header)
var headLock sync.Mutex
var head = &types.Header{Number: big.NewInt(0)}

func handleHeaders(opts *Options) {
	Log.Println("sub: awaiting new headers")
	for {
		header := <-headerChan
		hash := header.Hash().String()

		if _challengeHEAD(header) {
			continue
		}

		Log.Println("sub: new header recieved", hash)
		Log.Println("- fetching block for header...")
		go gather.Blocks(&gather.Options{
			Target:            bson.M{"hash": header.Hash().String()},
			Client:            opts.Client,
			DB:                opts.GatherOptions.DB,
			GetAncestors:      opts.GatherOptions.GetAncestors,
			SyncOptions:       opts.GatherOptions.SyncOptions,
			ShouldSynchronize: opts.GatherOptions.ShouldSynchronize,
		})
	}
}

func _challengeHEAD(header *types.Header) (isHead bool) {
	headLock.Lock()
	defer headLock.Unlock()

	// check if this was previous head
	hash := header.Hash().String()
	headHash := head.Hash().String()
	if hash == headHash {
		return true
	}

	// check if new head
	if header.Number.Uint64() > head.Number.Uint64() {
		head = header
	}

	Log.Println("sub: HEAD is", head.Number.Uint64(), headHash)
	return false
}
