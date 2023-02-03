package cmd

import (
	"crypto/x509"
	"fmt"
	"github.com/chramb/crtchkr/util"
	"github.com/chramb/crtchkr/util/action/discord"
	"github.com/chramb/crtchkr/util/action/mail"
	"github.com/chramb/crtchkr/util/failed"
	"github.com/urfave/cli/v2"
	"strings"
)

var shortFlag int = 0

var checkCmd = &cli.Command{
	Name:    "check",
	Usage:   "check validity of certificate",
	Aliases: []string{"c", "verify", "v"},
	Before:  checkAction,
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "oneline", Aliases: []string{"1"}, Usage: "Print shorter output", Count: &shortFlag},
		&cli.StringFlag{Name: "on-fail", Aliases: []string{"x"}, Usage: "Action to perform if certificate is invalid", Action: onFail},
	},
	Action: func(context *cli.Context) error {
		return nil
	},
}

func checkAction(ctx *cli.Context) error {
	argsNum := ctx.Args().Len()
	if argsNum < 1 {
		fmt.Println("Please provide URL or Path to a certificate")
	}
	for i := 0; i < argsNum; i++ {
		link := ctx.Args().Get(i)
		certs, err := util.GetCerts(link)
		if err != nil {
			return err
		}
		fmt.Printf("Found %d certificates: %s\n", len(certs), link)
		certsLen := len(certs)
		roots, _ := x509.SystemCertPool()
		for i := certsLen - 1; i >= 0; i-- {
			cert := certs[i]
			if shortFlag == 0 {
				fmt.Printf("\n---- %d. ORG: %s ----\n", i, cert.Subject.Organization)
			}
			// TODO: if invalid fail gracefully and print error why failed
			// TODO: check how many days left
			chains, err := util.CheckCert(certs[i], roots)
			if err != nil {
				failed.Fc = append(failed.Fc, failed.FailedCert{Cert: cert, Err: err})
				if shortFlag == 0 {
					fmt.Printf("%s invalid certificate: %s\n", cert.Subject.CommonName, err)
					fmt.Printf("signed by: %s\n", cert.Issuer.CommonName)
				} else {
					fmt.Printf("❌ - %s - %s - %s - Reason: %s\n", cert.Subject.Organization, cert.DNSNames, cert.Subject.CommonName, err)
				}
			}
			if chains != nil {
				if shortFlag == 0 {
					fmt.Printf("Valid Certificate! \nCommon Name: %s\n", cert.Subject.CommonName)
				} else {
					fmt.Printf("✅ - %s - %s - %s\n", cert.Subject.Organization, cert.DNSNames, cert.Subject.CommonName)
				}
				if cert.IsCA {
					roots.AddCert(cert)
					// DEBUG: println("adding ", cert.Subject.CommonName, " to roots")
				}
			}
		}
	}
	return nil
}

func onFail(ctx *cli.Context, arg string) error {
	if len(failed.Fc) == 0 {
		return nil
	}
	opts := strings.Split(arg, ".")
	switch opts[0] {
	case "discord":
		discord.Message(ctx, opts[1])
	case "mail":
		mail.Mail(ctx, opts[1])
	}
	return nil
}
