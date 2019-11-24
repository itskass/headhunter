package cmds

import (
	"io/ioutil"
	"log"

	"github.com/urfave/cli"
)

var Docs = cli.Command{
	Name:  "docs",
	Usage: "generate documentation for HEADHunter",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "out",
			Usage: "path to file to save output in.",
		},
	},
	Action: _docs,
}

func _docs(c *cli.Context) error {
	md, err := c.App.ToMarkdown()
	if err != nil {
		log.Fatal("err:", err)
	}

	ioutil.WriteFile(c.String("out"), []byte(md), 0700)
	return nil
}
