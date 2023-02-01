package cmd

import (
	"fmt"
	"github.com/chramb/crtchkr/util"
	"github.com/urfave/cli/v2"
)

// TODO: check all and success if any one is valid and has chain
var CheckCmd = &cli.Command{
	Name:  "check",
	Usage: "check validity of certificate",
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
			for _, cert := range certs {
				fmt.Printf("\n--Cert--\n")
				fmt.Printf("cn: %s", cert.Subject.CommonName)
				// TODO: if invalid fail gracefully and print error why failed
				// TODO: check how many days left
				chains, err := util.CheckCert(cert)
				_ = chains
				//fmt.Printf("chains: %s", chains)
				fmt.Printf("%s\n", err)
			}
		}
		return nil
	},
}
