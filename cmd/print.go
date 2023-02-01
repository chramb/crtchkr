package cmd

import (
	"fmt"
	"github.com/chramb/crtchkr/util"
	"github.com/urfave/cli/v2"
)

var printCmd = &cli.Command{
	Name:    "print",
	Usage:   "print certificate info",
	Aliases: []string{"p"},
	Action:  printCert,
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "dnsnames", Usage: "Print DNSNames specified in certificate", Count: &setFlags.dnsnames},
		&cli.BoolFlag{Name: "subject", Usage: "Print Subject specified in certificate", Count: &setFlags.subject},
		&cli.BoolFlag{Name: "issuer", Usage: "Print Issuer specified in certificate", Count: &setFlags.issuer},
		&cli.BoolFlag{Name: "version", Usage: "Print Certificate Version specified in certificate", Count: &setFlags.version},
		&cli.BoolFlag{Name: "not-before", Usage: "print date not before which certificate is valid", Count: &setFlags.not_before},
		&cli.BoolFlag{Name: "not-after", Usage: "print date when certificate expires", Count: &setFlags.not_after},
		&cli.BoolFlag{Name: "all", Aliases: []string{"a"}, Usage: "print all", Count: &setFlags.all},
	},
}

type printFlags struct {
	dnsnames   int
	subject    int
	issuer     int
	version    int
	not_before int
	not_after  int
	all        int
}

var setFlags = printFlags{
	dnsnames: 0, subject: 0, issuer: 0, version: 0, not_before: 0, not_after: 0, all: 0,
}

func printCert(ctx *cli.Context) error {
	argsNum := ctx.Args().Len()
	if argsNum < 1 {
		fmt.Println("Please provide URL or Path to .pem certificate file")
	}
	for i := 0; i < argsNum; i++ {
		link := ctx.Args().Get(i)
		certs, err := util.GetCerts(link)
		if err != nil {
			panic(err)
		}
		if argsNum > 1 {
			fmt.Printf("\n----- %d: %s -----\n", i, certs[0].Subject.CommonName)
		}
		for _, cert := range certs {
			if len(certs) > 1 {
				fmt.Printf("\n  --- %d: %s ---\n", i, cert.Subject.CommonName)
			}
			if setFlags.all > 0 || setFlags.subject > 0 {
				fmt.Println("  Subject: \t", cert.Subject)
			}
			if setFlags.all > 0 || setFlags.issuer > 0 {
				fmt.Println("  Issuer: \t", cert.Issuer)
			}
			if setFlags.all > 0 || setFlags.dnsnames > 0 {
				fmt.Println("  DNSNames: \t", cert.DNSNames)
			}
			if setFlags.all > 0 || setFlags.version > 0 {
				fmt.Println("  Version: \t", cert.Version)
			}
			if setFlags.all > 0 || setFlags.not_before > 0 {
				fmt.Println("  Not Before: \t", cert.NotBefore)
			}
			if setFlags.all > 0 || setFlags.not_after > 0 {
				fmt.Println("  Not After: \t", cert.NotAfter)
			}
		}
	}
	return nil
}
