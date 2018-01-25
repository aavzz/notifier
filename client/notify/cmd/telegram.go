package cmd

import (
	"crypto/tls"
	"encoding/json"
	"github.com/aavzz/misc/pipe"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var telegram = &cobra.Command{
	Use:   "telegram",
	Short: "Sends a message to telegram group",
	Long:  `Instructs notifyd to send message to telegram group`,
	Run:   telegramCommand,
}

func telegramCommand(cmd *cobra.Command, args []string) {

	type JResp struct {
		Error    int
		ErrorMsg string
	}

	//read message from stdin (pipe)
	message, err := pipe.Read(800)
	if err != nil {
		log.Fatal(err.Error())
	}

	params := url.Values{
		"channel":    {"telegram"},
		"recipients": {viper.GetString("telegram.group")},
		"message":    {message},
	}

	url := viper.GetString("telegram.url")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	resp, err := c.PostForm(url, params)
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		var v JResp
		if err := json.Unmarshal(body, &v); err != nil {
			log.Fatal(err)
		}

		if v.Error == 0 {
			log.Print("Message sent successfully")
		} else {
			log.Print(v.ErrorMsg)
		}
	} else {
		log.Print(resp.Status)
	}
}
