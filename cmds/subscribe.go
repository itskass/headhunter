package cmds

import (
	"io/ioutil"
	"time"

	"github.com/itskass/headhunter/subscribe"

	"github.com/itskass/headhunter/gather"
	"github.com/itskass/headhunter/helpers"
	"github.com/urfave/cli"
)

var Subscribe = cli.Command{
	Name:  "subscribe",
	Usage: "subscribe listens for and downloads the latests blocks",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "connect",
			Usage: "connects HEAD to a known ancestor, by downloading the missing blocks, should be true if delay is greater than average block time",
		},
		&cli.IntFlag{
			Name:  "delay",
			Usage: "delay in seconds between requests for latest block",
			Value: 30,
		},
		&cli.BoolFlag{
			Name:  "sync",
			Usage: "ensures you gather missing blocks",
		},
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "show all log outputs",
		},
		&cli.BoolFlag{
			Name:  "silent",
			Usage: "silences all logs",
		},
	},
	Action: _subscribe,
}

func _subscribe(c *cli.Context) error {

	// set logging options
	if !c.Bool("verbose") {
		gather.Log.SetFlags(0)
		gather.Log.SetOutput(ioutil.Discard)
	}
	if c.Bool("silent") {
		gather.Log.SetFlags(0)
		gather.Log.SetOutput(ioutil.Discard)
		gather.SyncLog.SetFlags(0)
		gather.SyncLog.SetOutput(ioutil.Discard)
		subscribe.Log.SetFlags(0)
		subscribe.Log.SetOutput(ioutil.Discard)
	}

	// create gather options
	opts := &gather.Options{
		DB:                helpers.DB(c.String("db")),
		Client:            helpers.Client(c.String("rpc")),
		GetAncestors:      c.Bool("ancestors"),
		ShouldSynchronize: c.Bool("sync"),
	}

	// create sub options
	subOpts := &subscribe.Options{
		Client:        opts.Client,
		Delay:         time.Second * time.Duration(c.Int("delay")),
		GatherOptions: opts,
	}

	// run
	subscribe.HTTP(subOpts)
	return nil
}
