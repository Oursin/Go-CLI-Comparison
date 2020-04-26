package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"test-urfave-cli/cmd"
)

func main() {
	app := &cli.App{
		Name: "c3pm",
		Usage: "c3pm usage",
		Commands: cmd.Commands,
	}

	err := app.Run(os.Args)
	if (err != nil) {
		log.Fatal(err)
	}
}