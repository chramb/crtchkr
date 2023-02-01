package cmd

import (
	"github.com/chramb/crtchkr/util"
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:    "crtchkr",
	Usage:   "check validity of local and remote x509 certificates",
	Version: "0.0.1-pre",
	Action:  cli.ShowAppHelp,
	Flags:   []cli.Flag{},
	Commands: []*cli.Command{
		printCmd,
		checkCmd,
		msgCmd,
	},
}

var VersionFlag = &cli.BoolFlag{
	Name:               "version",
	Aliases:            []string{"V"},
	Usage:              "print version",
	DisableDefaultText: true,
}

func raw(ctx *cli.Context) error {
	argsNum := ctx.Args().Len()
	if argsNum < 1 {
		// TODO: read certs from config
		println("TODO: Get certs from toml file")
		return nil
	}
	for i := 0; i < argsNum; i++ {
		link := ctx.Args().Get(i)
		util.GetCerts(link)
	}
	return nil
}
