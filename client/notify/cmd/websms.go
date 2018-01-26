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

var websms = &cobra.Command{
	Use:   "websms",
	Short: "Sends an SMS via websms",
	Long:  `Instructs notifyd to send SMS via websms`,
	Run:   websmsCommand,
}

func websmsCommand(cmd *cobra.Command, args []string) {

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
		"channel":    {"websms"},
		"recipients": {viper.GetString("websms.recipients")},
		"message":    {message},
	}

	url := viper.GetString("websms.url")

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
