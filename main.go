package main

import (
	"log"
	"time"

	"github.com/itskass/headhunter/cmds"

	"github.com/urfave/cli"

	"github.com/itskass/headhunter/subscribe"

	"github.com/itskass/headhunter/gather"

	"github.com/globalsign/mgo"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	app := cli.NewApp()
	app.Name = "HEADHunter"
	app.Version = "0.1"
	app.Author = "kassius Barker"
	app.Email = "itskass94@gmail.com"
	app.Description =
		`HEADHunter is a tool for downloading/syncing ethereum-like blockchains to MongoDB via a nodes rpc.
		This software also provides some tools, quering blocks in the database.`

	// global flags
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "db",
			Usage: "url of running MongoDB instance",
			Value: "localhost:27017",
		},
		&cli.StringFlag{
			Name:     "rpc",
			Usage:    "url of blockchain ethereum rpc",
			Required: true,
		},
	}

	// register commands
	app.Commands = []cli.Command{
		cmds.Gather,
		cmds.Subscribe,
		cmds.Docs,
	}

	// run application
	app.RunAndExitOnError()
}

func oldmain() {

	log.Println("db: opening connection to mongo...")
	sess, err := mgo.Dial("localhost:27017")
	handle(err)

	log.Println("db: conencted")
	db := sess.DB("blockchain")

	log.Println("rpc: dailing nodes rpc...")
	client, err := ethclient.Dial("http://51.83.46.243:8545")
	handle(err)

	log.Println("rpc: connected")

	// log.Println("getting block")
	// gather.Blocks(&gather.Options{
	// 	Target: bson.M{"number": uint64(10)},

	// 	Client:       client,
	// 	DB:           db,
	// 	GetAncestors: false,
	// })

	gatherOpts := &gather.Options{
		Client:       client,
		DB:           db,
		GetAncestors: true,
	}

	subscribeOpts := &subscribe.Options{
		Client:        client,
		GatherOptions: gatherOpts,
		Delay:         time.Second * 15,
	}

	subscribe.HTTP(subscribeOpts)
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
