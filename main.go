package main

import (
	"crtchkr/cmd"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = cmd.VersionFlag
	app := cmd.App

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
