package discord

import (
	"fmt"
	"net/http"
	"strings"
)

type Discord struct {
	Url string
}

func (d *Discord) Post(cn string, error string) {
	data := fmt.Sprintf(`
	{
		"content": "%s"
	}
	`, error)

	client := &http.Client{}
	myReader := strings.NewReader(data)
	req, err := http.NewRequest("POST", d.Url, myReader)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	fmt.Println(resp)

}
