package cmds

import (
	"io/ioutil"

	"github.com/globalsign/mgo/bson"
	"github.com/itskass/headhunter/gather"
	"github.com/itskass/headhunter/helpers"
	"github.com/urfave/cli"
)

var Gather = cli.Command{
	Name:  "gather",
	Usage: "gather downloads specified the latests blocks",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:  "number",
			Usage: "target block by number (index)",
		},
		&cli.StringFlag{
			Name:  "hash",
			Usage: "target block via hash",
		},
		&cli.BoolFlag{
			Name:  "connect",
			Usage: "connects target to known ancestor, by downloading the missing blocks",
		},
		&cli.BoolFlag{
			Name:  "sync",
			Usage: "gather missing blocks",
		},
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "show all log outputs when synchronizing",
		},
		&cli.BoolFlag{
			Name:  "silent",
			Usage: "silences all logs",
		},
	},
	Action: _gather,
}

func _gather(c *cli.Context) error {

	// set logging options
	if !c.Bool("verbose") && c.Bool("sync") {
		gather.Log.SetFlags(0)
		gather.Log.SetOutput(ioutil.Discard)
	}
	if c.Bool("silent") {
		gather.Log.SetFlags(0)
		gather.Log.SetOutput(ioutil.Discard)
		gather.SyncLog.SetFlags(0)
		gather.SyncLog.SetOutput(ioutil.Discard)
	}

	// create gather options
	opts := &gather.Options{
		DB:                helpers.DB(c.GlobalString("db")),
		Client:            helpers.Client(c.GlobalString("rpc")),
		GetAncestors:      c.Bool("connect"),
		ShouldSynchronize: c.Bool("sync"),
	}

	// gather target
	if c.IsSet("number") {
		opts.Target = bson.M{"number": c.Uint64("number")}
	} else if c.IsSet("hash") {
		opts.Target = bson.M{"hash": c.String("hash")}
	} else {
		opts.Target = bson.M{"latest": "true"}
	}

	// run
	gather.Blocks(opts)
	return nil
}
