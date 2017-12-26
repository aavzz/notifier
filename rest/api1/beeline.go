package api1

import (
	"fmt"
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"encoding/xml"
	"golang.org/x/text/encoding/charmap"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/cfgfile"
	
)

func sendMessageBeeline(numbers string, message string) {
  
	type Output struct {
		Errors []string `xml:"errors>error"`
	}
	
	cfg, err := CfgFileContent()
	if err != nil {
		SysLog.Err(err.Error())	
	} else {
		msg, err := charmap.Windows1251.NewEncoder().String(message)
		if err != nil {
			SysLog.Err(err.Error())	
		} else {
			parameters := url.Values{
				"user": {cfg.Beeline.Login},
				"pass": {cfg.Beeline.Password},
				"sender": {cfg.Beeline.Sender},
				"action": {"post_sms"},
				"target": {numbers},
				"message": {msg},
			}

			url := "https://beeline.amega-inform.ru/sendsms/"
			req, err := http.NewRequest("POST", url, strings.NewReader(parameters.Encode())) 
			if err != nil {
				SysLog.Err(err.Error())
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=windows-1251")
	
			tr := &http.Transport{
		        	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			c := &http.Client{Transport: tr}

		  	resp, err := c.Do(req)
			if err != nil {
				SysLog.Err(err.Error())
			}
			if resp != nil {
				defer resp.Body.Close()
			}
			
			var v Output;
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				SysLog.Err(err.Error())
			}
			err = xml.Unmarshal(body, &v)
			if err != nil {
				SysLog.Err(err.Error())
			}
			
			if v.Errors != nil {
				SysLog.Err(fmt.Sprintf("Provider output: %q", v.Errors))	
			}
		}
	}
}
