package subscribe

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
)

// WS will subscribe to "newHeads" using websockets
// and gather/save new blocks .
func WS(opts *Options) (ethereum.Subscription, error) {

	log.Println("sub: attempting to subscribe to \"newHeads\"")
	sub, err := opts.Client.SubscribeNewHead(context.Background(), headerChan)
	if err != nil {
		log.Println("sub: err:", err)
		return nil, err
	}

	log.Println("sub: subscribed to \"newHeads\"")
	go handleHeaders(opts)
	return sub, nil
}
