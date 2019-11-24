package helpers

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/globalsign/mgo"
)

func DB(url string) *mgo.Database {
	sess, err := mgo.Dial(url)
	if err != nil {
		log.Println("db: FATAL: couldn't connect:")
		log.Fatal("err:", err)
	}
	log.Println("db: connected")
	return sess.DB("blockchain")
}

func Client(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Println("rpc: FATAL: couldn't connect")
		log.Fatal("err:", err)
	}
	log.Println("rpc: connected")
	return client
}
