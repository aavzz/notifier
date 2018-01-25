package channels

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"regexp"
	"golang.org/x/text/encoding/charmap"
	"strings"
)

// BeelineResponse holds the response from beeline
// Must be exportable
type BeelineResponse struct {
	Errors []string `xml:"errors>error"`
}

// SendMessageBeeline sends message via beeline
func SendMessageBeeline(login, pass, sender, recipients, msg string) error {

	msg, err := charmap.Windows1251.NewEncoder().String(msg)
	if err != nil {
		return err
	} else {
                recipients := strings.Join(regexp.MustCompile(`\+\d+`).FindAllString(recipients, -1), ",")

		l := len(msg)
		if l > 480 {
			l = 480
		}
		msg := msg[:l]

		params := url.Values{
			"user":    {login},
			"pass":    {pass},
			"sender":  {sender},
			"action":  {"post_sms"},
			"target":  {recipients},
			"message": {msg},
		}

		c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		resp, err := c.PostForm("https://beeline.amega-inform.ru/sendsms/", params)
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
			var v BeelineResponse
			err = xml.Unmarshal(body, &v)
			if err != nil {
				return err
			}
			if v.Errors != nil {
				return errors.New(fmt.Sprintf("Provider output: %q", v.Errors))
			}
		} else {
			return errors.New(fmt.Sprintf("Provider output: %s", resp.Status))
		}
	}
	return nil
}
