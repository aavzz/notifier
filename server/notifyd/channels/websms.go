package channels

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// WebsmsResponse holds the response from smsc
// Must be exportable
type WebsmsResponse struct {
	Error string
}

// SendMessageWebsms sends message via websms
func SendMessageWebsms(login, pass, sender, recipients, msg string) error {

	msg, err := charmap.Windows1251.NewEncoder().String(msg)
	if err != nil {
		return err
	} else {
		recipients := strings.Join(regexp.MustCompile(`\+\d+`).FindAllString(recipients, -1), ",")

		l := len(msg)
		if l > 800 {
			l = 800
		}
		msg := msg[:l]

		params := url.Values{
			"json": {"{\"http_username\": \"" + login + "\", \"http_password\": \"" + pass + "\", \"message\": \"" + msg + "\", \"phone_list\": \"" + recipients + "\", \"fromPhone\": \"" + sender + "\"}"},
		}

		c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		resp, err := c.PostForm("https://cab.websms.ru/json_in5.asp", params)
		if err != nil {
			return err
		}
		if resp != nil {
			defer resp.Body.Close()
		}
		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			var v WebsmsResponse
			err = json.Unmarshal(body, &v)
			if err != nil {
				return err
			}

			if v.Error != "OK" {
				return errors.New(fmt.Sprintf("Provider output: %q", v.Error))
			}
		} else {
			return errors.New(fmt.Sprintf("Provider output: %s", resp.Status))
		}
	}
	return nil
}
