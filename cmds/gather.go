package cmds

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/urfave/cli"

	"github.com/globalsign/mgo/bson"
	"github.com/itskass/headhunter/gather"
	"github.com/itskass/headhunter/helpers"
)

var Gather = cli.Command{
	Name:  "gather",
	Usage: "gather downloads specified the blocks. If no blocks are specified the latest block is used as a target",
	Flags: []cli.Flag{
		&cli.Uint64Flag{
			Name:  "number",
			Usage: "target block by number (index)",
		},
		&cli.StringFlag{
			Name:  "range",
			Usage: "target blocks in the given range <start:end>",
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
		DB:                helpers.DB(c.String("db")),
		Client:            helpers.Client(c.String("rpc")),
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

	// run range
	if c.IsSet("range") {
		log.Println("downloading range:")
		start, end := parseRange(c.String("range"))
		for i := start; i < end; i++ {
			opts.Target = bson.M{"number": uint64(i)}
			opts.GetAncestors = false      // cant get ancestors for range
			opts.ShouldSynchronize = false // cant synchronize for rane
			gather.Blocks(opts)
		}
		return nil
	}

	// run default
	gather.Blocks(opts)
	return nil
}

func parseRange(str string) (int, int) {
	s := strings.Split(str, ":")
	if len(s) != 2 {
		log.Fatal("bad range:", str)
	}
	start, _ := strconv.Atoi(s[0])
	end, _ := strconv.Atoi(s[1])
	return start, end
}
