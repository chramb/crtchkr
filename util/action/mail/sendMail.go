package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/chramb/crtchkr/util"
	"github.com/chramb/crtchkr/util/failed"
	"github.com/urfave/cli/v2"
	"net/smtp"
	"text/template"
)

func Mail(ctx *cli.Context, key string) (err error) {
	config, _, err := util.GetConfig(ctx.String("config"))
	if err != nil {
		fmt.Println(err)
	}
	mailConfig := config.Mail[key]
	if mailConfig.From == "" {
		fmt.Println(config.Mail[key])
		return cli.Exit("Specified table doesn't exist in config", -1)
	}
	from := mailConfig.From
	to := mailConfig.To
	t, err := template.New("mail_content").Parse(config.Mail[key].Message)
	if err != nil {
		fmt.Println(err)
	}
	var buf bytes.Buffer
	t.Execute(&buf, failed.Fc)

	a := mailConfig.Auth
	auth := smtp.PlainAuth(a.Identity, a.Username, a.Password, a.Host)

	// Connect to the Gmail SMTP server using StartTLS
	conn, err := tls.Dial("tcp", mailConfig.Server, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, "smtp.gmail.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	// Authenticate to the Gmail SMTP server
	if err = client.Auth(auth); err != nil {
		fmt.Println(err)
		return
	}

	// Send the email
	if err = client.Mail(from); err != nil {
		fmt.Println(err)
		return
	}
	for _, recipient := range to {
		if err = client.Rcpt(recipient); err != nil {
			fmt.Println(err)
			return
		}
	}
	w, err := client.Data()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	err = w.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email sent successfully.")
	return nil
}
