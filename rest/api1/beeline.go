package api1

import (
	"fmt"
	"bytes"
	"net/url"
	"github.com/gorilla/http"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/cfgfile"
)

func sendMessageBeeline(numbers string, message string) {
  
	headers := make(map[string][]string)
	headers["Content-Type"][0] = "application/x-www-form-urlencoded; charset=windows-1251"
	
	parameters := url.Values{}
	parameters.Add("user", ConfigBeelineLogin())
	parameters.Add("pass", ConfigBeelinePassword())
	parameters.Add("sender", ConfigBeelineSender())
	parameters.Add("action", "post_sms")
	parameters.Add("target", numbers)
	parameters.Add("message", message)
  
  	c := new(http.Client)

	status, _, r, err := c.Post("https://beeline.amega-inform.ru/sendsms/", headers, bytes.NewBufferString(parameters.Encode()))
	if err != nil {
		SysLog.Err(err.Error())
	}
	if r != nil {
		defer r.Close()
	}
	SysLog.Info(fmt.Sprintf("Post result: %v", status))
}
