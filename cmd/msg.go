package cmd

import (
	"github.com/chramb/crtchkr/util/action/discord"
	"github.com/chramb/crtchkr/util/action/mail"
	"github.com/urfave/cli/v2"
)

var msgCmd = &cli.Command{
	Name: "message",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "Discord", Aliases: []string{"d"}, Action: discord.Message},
		&cli.StringFlag{Name: "Mail", Aliases: []string{"m"}, Action: mail.Mail},
	},
	Action: cli.ShowSubcommandHelp,
}
