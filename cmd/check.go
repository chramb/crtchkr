package cmd

import (
	"crypto/x509"
	"fmt"
	"github.com/chramb/crtchkr/util"
	"github.com/urfave/cli/v2"
)

// TODO: check all and success if any one is valid and has chain
var checkCmd = &cli.Command{
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
			fmt.Printf("Found %d certificates.\n", len(certs))
			certsLen := len(certs)
			roots, _ := x509.SystemCertPool()
			for i := certsLen - 1; i >= 0; i-- {
				fmt.Printf("\n---- %d. ----\n", i)
				cert := certs[i]
				// TODO: if invalid fail gracefully and print error why failed
				// TODO: check how many days left
				chains, err := util.CheckCert(certs[i], roots)
				if err != nil {
					fmt.Printf("%s invalid certificate: %s\n", cert.Subject.CommonName, err)
					fmt.Printf("signed by: %s\n", cert.Issuer.CommonName)
					if i != 0 {
						return cli.Exit("Warning Broken Chain!!", -666)
					}
				}
				if chains != nil {
					fmt.Printf("Valid Certificate! %s\n", cert.Subject.CommonName)
					fmt.Printf("Chain lenght %d\n", len(chains))
					if cert.IsCA {
						roots.AddCert(cert)
						// DEBUG: println("adding ", cert.Subject.CommonName, " to roots")
					}
				}
			}
		}
		return nil
	},
}
