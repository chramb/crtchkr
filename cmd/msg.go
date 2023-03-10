package cmd

import (
	"github.com/chramb/crtchkr/util/action/discord"
	"github.com/chramb/crtchkr/util/action/mail"
	"github.com/urfave/cli/v2"
)

var msgCmd = &cli.Command{
	Name:    "message",
	Usage:   "execute message action",
	Aliases: []string{"m"},
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "discord", Aliases: []string{"d"}, Action: discord.Message, Usage: "run discord action from provided config [discord.`VALUE`]"},
		&cli.StringFlag{Name: "Mail", Aliases: []string{"m"}, Action: mail.Mail, Usage: "run mail action from provided config [mail.`VALUE`]"},
	},
	Action: func(context *cli.Context) error {
		return nil
	},
}
