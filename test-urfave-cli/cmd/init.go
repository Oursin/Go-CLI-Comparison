package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "init a new c3pm project",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "name",
			Aliases: []string{"n"},
			Usage: "name of the project to init",
		},
	},
	Action: func(c *cli.Context) error {
		fmt.Printf("Will init %s\n", c.String("name"))
		return nil
	},
}
