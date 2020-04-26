# Go-CLI-Comparison

This repository hosts several dummy projects made in order to try out Go CLI libraries.

## Libraries considered

The library tested are:

- [urfave/cli](https://github.com/urfave/cli) v2 ([documentation](https://pkg.go.dev/github.com/urfave/cli/v2?tab=doc))
- [cobra](https://github.com/spf13/cobra) ([documentation](https://pkg.go.dev/github.com/spf13/cobra?tab=doc))
- [ffcli](https://github.com/peterbourgon/ff) v3 ([documentation](https://pkg.go.dev/github.com/peterbourgon/ff/v3/ffcli?tab=doc))
- [kong](https://github.com/alecthomas/kong) ([documentation](https://pkg.go.dev/github.com/alecthomas/kong?tab=doc))

## Test specification

The specs of the test are as follows:

- The CLI should be able to handle flags easily
- The CLI should be able to handle flag aliases
- The CLI should be able to handle subcommands (such as `git add`) 
