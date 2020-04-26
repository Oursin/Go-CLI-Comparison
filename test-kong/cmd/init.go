package cmd

import "fmt"

type InitCmd struct {
	Name string `help:"The name of the project to init"`
}

func (i *InitCmd) Run(ctx *Context) error {
	fmt.Printf("Will init %s\n", i.Name)
	return nil
}