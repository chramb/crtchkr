package cmd

import (
	"github.com/urfave/cli/v2"
)

var configFlag = &cli.StringFlag{
	Name:    "config",
	Aliases: []string{"c"},
	Usage:   "Load configuration from specified `FILE`",
	//EnvVars:
}
