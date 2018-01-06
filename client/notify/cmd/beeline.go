package cmd

import (
	"crypto/tls"
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var beeline = &cobra.Command{
	Use:   "notify beeline",
	Short: "sends SMS via beeline",
	Long:  `Instructs notifyd to send SMS via beeline`,
	Run:   beelineCommand,
}

func beelineCommand(cmd *cobra.Command, args []string) {

	type JResp struct {
		Error    int
		ErrorMsg string
	}

	parameters := url.Values{
		"channel": {"beeline"},
		"recipients": {viper.GetString("beeline.recipients")},
		"message":     {viper.GetString("beeline.message")},
	}

	url := viper.GetString("beeline.url")
	req, err := http.NewRequest("POST", url, strings.NewReader(parameters.Encode()))
	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	if resp != nil {
		defer resp.Body.Close()
	}

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
}

