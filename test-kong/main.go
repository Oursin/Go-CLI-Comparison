package main

import (
	"github.com/alecthomas/kong"
	"test-kong/cmd"
)

func main() {
	ctx := kong.Parse(&cmd.CLI)
	err := ctx.Run(&cmd.Context{Debug: cmd.CLI.Debug})
	ctx.FatalIfErrorf(err)
}