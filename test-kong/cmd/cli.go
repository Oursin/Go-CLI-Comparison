package cmd

type Context struct {
	Debug bool
}

var CLI struct{
	Debug bool `help:"Enable debug mode"`

	Init InitCmd `cmd help:"Init a c3pm project"`
}
