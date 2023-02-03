package discord

import (
	"bytes"
	"fmt"
	"github.com/chramb/crtchkr/util"
	"github.com/chramb/crtchkr/util/failed"
	"github.com/urfave/cli/v2"
	"net/http"
	"strings"
	"text/template"
)

func Message(ctx *cli.Context, key string) error {
	config, _, err := util.GetConfig(ctx.String("config"))
	t, err := template.New("webhook_data").Parse(config.Discord[key].Request)
	if err != nil {
		fmt.Println(err)
	}
	var buf bytes.Buffer
	t.Execute(&buf, failed.Fc)
	data := buf.String()

	client := &http.Client{}
	myReader := strings.NewReader(data)
	req, err := http.NewRequest("POST", config.Discord[key].Url, myReader)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Webhook Sent")
	return nil
}
