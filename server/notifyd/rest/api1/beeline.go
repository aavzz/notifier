package api1

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"errors"
)

func sendMessageBeeline(numbers string, message string) error {

	// Must be exportable
	type Output struct {
		Errors []string `xml:"errors>error"`
	}

	msg, err := charmap.Windows1251.NewEncoder().String(message)
	if err != nil {
		return err
	} else {
		parameters := url.Values{
			"user":    {viper.GetString("beeline.Login")},
			"pass":    {viper.GetString("beeline.Password")},
			"sender":  {viper.GetString("beeline.Sender")},
			"action":  {"post_sms"},
			"target":  {numbers},
			"message": {msg},
		}

		url := "https://beeline.amega-inform.ru/sendsms/"
		req, err := http.NewRequest("POST", url, strings.NewReader(parameters.Encode()))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=windows-1251")

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c := &http.Client{Transport: tr}

		resp, err := c.Do(req)
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
			var v Output
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
