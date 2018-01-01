/*
Package api1 implements version 1 of notifyd API.
*/
package api1

import (
	"fmt"
	. "github.com/aavzz/notifier/setup/syslog"
	"net/http"
	"regexp"
	"strings"
)

// Handler calls the right function to send message via specified channel.
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

		senderName := r.FormValue("sender_name")
		senderAddr := r.FormValue("sender_address")
		subject := r.FormValue("subject")

		re := regexp.MustCompile(`\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`)
		emails := re.FindAllString(recipients, 5)
		senderAddress := re.FindAllString(senderAddr, 1)
		l := len(message)
		if l > 480 {
			l = 480
		}
		msg := message[:l]
		if emails != nil && msg != "" {
			SysLog.Info(fmt.Sprintf("Message '%s' sent via email to %q", msg, emails))
			sendMessageEmail(senderName, senderAddress, emails, subject, msg)
		} else {
			SysLog.Info(fmt.Sprintf("Failed to send message via email"))
		}
	default:
		SysLog.Info("No valid channel found")
	}
}
