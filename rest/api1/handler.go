package api1

import (
	"fmt"
	"regexp"
	"strings"
	"net/http"
	. "github.com/aavzz/notifier/setup/syslog"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	channel := r.FormValue("channel")
	recipients := r.FormValue("recipients")
	message := r.FormValue("message")

	switch channel {
		case "beeline":
			re := regexp.MustCompile(`\+7\d{10}`)
			phones := strings.Join(re.FindAllString(recipients, 5), ",")
			l := len(message)
			if l > 480 {
				l = 480
			}
			msg := message[:l]
			if phones != "" && msg != "" {
				SysLog.Info(fmt.Sprintf("Message '%s' sent via beeline to %s", msg, phones))
				sendMessageBeeline(phones, msg)
			} else {
				SysLog.Info(fmt.Sprintf("Failed to send message via beeline"))
			}
		case "email":
			re := regexp.MustCompile(`\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`)
			emails := re.FindAllString(recipients, 5)
			l := len(message)
			if l > 480 {
				l = 480
			}
			msg := message[:l]
			if emails != nil && msg != "" {
				SysLog.Info(fmt.Sprintf("Message '%s' sent via email to %q", msg, emails))
				sendMessageEmail(emails, msg)
			} else {
				SysLog.Info(fmt.Sprintf("Failed to send message via email"))
			}
		
		default:
			SysLog.Info("No valid channel found")	
	}
}
