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
			SysLog.Info("BBBBBBB")
			phones := strings.Join(re.FindAllString(recipients, 5), ",")
			l := len(message)
			if l > 480 {
				l = 480
			}
			msg := message[:l]
			SysLog.Info("CCCC")
			SysLog.Info(phones)
			if phones != "" && msg != "" {
				sendMessageBeeline(phones, msg)
				SysLog.Info(fmt.Sprintf("Message '%s' sent via beeline to %s", msg, phones))
			} else {
				SysLog.Info(fmt.Sprintf("Failed to send message via beeline"))
			}
		case "email":
		
		case "telegram":

		default:
			SysLog.Info("No valid channel found")	
	}
		
}
