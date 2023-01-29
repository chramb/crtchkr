package cmd

import (
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:     "crtchkr",
	Usage:    "check validity of local and remote x509 certificates",
	Version:  "0.0.1-pre",
	Action:   cli.ShowAppHelp,
	Flags:    []cli.Flag{},
	Commands: []*cli.Command{},
}

var VersionFlag = &cli.BoolFlag{
	Name:               "version",
	Aliases:            []string{"V"},
	Usage:              "print version",
	DisableDefaultText: true,
}
