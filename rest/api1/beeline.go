package api1

import (
	"fmt"
	"strings"
	"net/url"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/cfgfile"
)

func sendMessageBeeline(numbers string, message string) {
  
	parameters := url.Values{
		"user": {ConfigBeelineLogin()},
		"pass": {ConfigBeelinePassword()},
		"sender": {ConfigBeelineSender()},
		"action": {"post_sms"},
		"target": {numbers},
		"message": {message},
	}
  
	req, err := http.NewRequest("POST", "https://beeline.amega-inform.ru/sendsms/", strings.NewReader(parameters.Encode())) 
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
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	SysLog.Info(fmt.Sprintf("Post result: %v", resp.Status))
	SysLog.Info(fmt.Sprintf("Post result: %v", bodyString))
}
