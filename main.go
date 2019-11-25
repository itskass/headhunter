package main

import (
	"github.com/itskass/headhunter/cmds"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "HEADHunter"
	app.Version = "0.1"
	app.Authors = []*cli.Author{
		{
			Name:  "kassius Barker",
			Email: "itskass94@gmail.com",
		},
	}
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
	app.Commands = []*cli.Command{
		&cmds.Gather,
		&cmds.Subscribe,
		&cmds.Docs,
	}

	// run application
	app.RunAndExitOnError()
}
