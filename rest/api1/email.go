package api1

import (
	"gopkg.in/gomail.v2"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/cfgfile"
)

func sendMessageEmail(emails []string, message string) {
 
	cfg, err := CfgFileContent()
	if err != nil {
		SysLog.Err(err.Error())	
	} else {
		//msg, err := charmap.Windows1251.NewEncoder().String(message)
		m := gomail.NewMessage()
		m.SetHeaders(map[string][]string{
			"From":    {m.FormatAddress("noreply@telinxet.ru", "Notifier")},
			"To":      emails,
			"Subject": {"** Notification [" + cfg.Email.Sender + "] **"},
		})
		m.SetBody("text/plain", message)
  
		d := gomail.Dialer{Host: "localhost", Port: 25}
		if err := d.DialAndSend(m); err != nil {
			SysLog.Err(err.Error())
		}
	}
}
