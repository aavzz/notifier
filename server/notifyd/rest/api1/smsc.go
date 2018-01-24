package api1

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func sendMessageSmsc(numbers string, message string) error {

	// Must be exportable
	type Output struct {
		error string
	}

	msg, err := charmap.Windows1251.NewEncoder().String(message)
	if err != nil {
		return err
	} else {
		parameters := url.Values{
			"login":  {viper.GetString("smsc.Login")},
			"psw":    {viper.GetString("smsc.Password")},
			"phones": {numbers},
			"mes":    {msg},
			"fmt":    {"3"},
		}

		url := "https://smsc.ru/sys/send.php"
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
			err = json.Unmarshal(body, &v)
			if err != nil {
				return err
			}

			if v.error != "" {
				return errors.New(fmt.Sprintf("Provider output: %q", v.error))
			}
		} else {
			return errors.New(fmt.Sprintf("Provider output: %s", resp.Status))
		}
	}
	return nil
}
