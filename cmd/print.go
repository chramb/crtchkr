package cmd

import (
	"fmt"
	"github.com/chramb/crtchkr/util"
	"github.com/urfave/cli/v2"
)

var printCmd = &cli.Command{
	Name:  "print",
	Usage: "print certificate info",
	Action: func(ctx *cli.Context) error {
		argsNum := ctx.Args().Len()
		if argsNum < 1 {
			fmt.Println("Please provide URL or Path to a certificate")
		}
		for i := 0; i < argsNum; i++ {
			link := ctx.Args().Get(i)
			certs, err := util.GetCerts(link)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Cert %d\n", i)
			for _, cert := range certs {
				fmt.Printf("%s: %s\n", cert.Subject.CommonName, cert.DNSNames)
			}
		}
		return nil
	},
}
