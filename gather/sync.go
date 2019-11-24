package gather

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

type SyncOptions struct {
	Head uint64
	Tail uint64
}

func Synchronize(opts *Options) {

	SyncLog.Println("sync: synchronizing")

	_head := opts.SyncOptions.Head
	_tail := opts.SyncOptions.Tail
	_targets := []uint64{}

	// build list of missing blocks between head and tail.
	for i := _tail; i < _head; i++ {
		if !owned(bson.M{"header.number": uint64(i)}, opts.DB) {
			_targets = append(_targets, i)
		}
	}

	// gather missing blocks
	SyncLog.Println("sync: missing blocks", len(_targets))
	for i := 0; i < len(_targets); i++ {
		if i%10 == 0 {
			log.Printf("sync: gathering missing blocks %d/%d", i, len(_targets))
		}

		// gather and save block
		Blocks(&Options{
			Target:       bson.M{"number": _targets[i]},
			Client:       opts.Client,
			DB:           opts.DB,
			GetAncestors: false,
		})
	}

	SyncLog.Println("sync: done")
}
