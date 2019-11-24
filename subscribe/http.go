package subscribe

import (
	"context"
	"log"
	"time"
)

func HTTP(opts *Options) {
	go handleHeaders(opts)
	for {
		getLatest(opts)
		time.Sleep(opts.Delay)
	}
}

func getLatest(opts *Options) {
	Log.Println("rpc: requesting latest header")
	// get latest block
	latest, err := opts.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Println("rpc: failed to get latest block")
		log.Println("rpc: err: ", err)
		return
	}
	// push latest block to channel
	headerChan <- latest
}
