package cmd

import (
	"github.com/chramb/crtchkr/util/action/discord"
	"github.com/chramb/crtchkr/util/action/exec"
	"github.com/urfave/cli/v2"
)

var msgCmd = &cli.Command{
	Name: "message",
	Action: func(ctx *cli.Context) error {

		_ = discord.Discord{
			Url: "https://discord.com/api/webhooks/1070338292601593866/qE8fYrlacMsE7r6430QKyxRnnbHy1uHqSuCDJp5pHV6lLHl44nMfeMzHEB4viSEr4Zfm",
		}
		//x.Post("test", "Your certificate Expired!!!")
		exec.Run("test", "Yooo")

		return nil
	},
}
