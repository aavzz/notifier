package channels

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"golang.org/x/text/encoding/charmap"
        "strings"
	"regexp"
)

// SmscResponse holds the response from smsc
// Must be exportable
type SmscResponse struct {
	Error string
}

// SendMessageSmsc sends message via smsc
func SendMessageSmsc(login, pass, sender, recipients, msg string) error {

        msg, err := charmap.Windows1251.NewEncoder().String(msg)
        if err != nil {
                return err
        } else {
                recipients := strings.Join(regexp.MustCompile(`\+\d+`).FindAllString(recipients), ",")

                l := len(msg)
                if l > 800 {
                        l = 800
                }
                msg := msg[:l]

		params := url.Values{
			"login":  {login},
			"psw":    {pass},
			"phones": {recipients},
			"sender": {sender},
			"mes":    {msg},
			"fmt":    {"3"},
		}

		c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
		resp, err := c.PostForm("https://smsc.ru/sys/send.php", params)
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
			var v SmscResponse
			err = json.Unmarshal(body, &v)
			if err != nil {
				return err
			}

			if v.Error != "" {
				return errors.New(fmt.Sprintf("Provider output: %q", v.Error))
			}
		} else {
			return errors.New(fmt.Sprintf("Provider output: %s", resp.Status))
		}
	}
	return nil
}
