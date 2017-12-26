package api1

import (
  "net/smtp"
  "gopkg.in/gomail.v2"
  . "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/cfgfile"
)

func sendMessageEmail(emails []string, message string) {
 
  m := gomail.NewMessage()
  m.SetHeaders(map[string][]string{
    "From":    {m.FormatAddress("alex@example.com", "Alex")},
    "To":      emails,
    "Subject": {"Hello"},
  })
  m.SetBody("text/plain", "Hello!")
  
  d := gomail.Dialer{Host: "localhost", Port: 25}
  if err := d.DialAndSend(m); err != nil {
    SysLog.Err(err.Error())
  }
  
}
