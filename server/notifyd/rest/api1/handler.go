/*
Package api1 implements version 1 of notifyd API.
*/
package api1

import (
	"encoding/json"
	"github.com/aavzz/daemon/log"
	"net/http"
	"regexp"
	"strings"
)

// Handler calls the right function to send message via specified channel.
func Handler(w http.ResponseWriter, r *http.Request) {

	//Must be exportable
	type JResponse struct {
		Error    int
		ErrorMsg string
	}

	var resp JResponse
	ret := json.NewEncoder(w)

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
			if err := sendMessageBeeline(phones, msg); err == nil {
				resp.Error = 0
				resp.ErrorMsg = "Message" + msg + "sent via" + channel + "to" + phones
				if err := ret.Encode(resp); err != nil {
					log.Error(err.Error())
				}
			} else {
				resp.Error = 1
                                resp.ErrorMsg = err.Error()
                                if err := ret.Encode(resp); err != nil {
                                        log.Error(err.Error())
                                }	
			}
		} else {
			resp.Error = 1
			resp.ErrorMsg = "Failed to send message via" + channel
			if err := ret.Encode(resp); err != nil {
				log.Error(err.Error())
			}
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
			if err := sendMessageEmail(senderName, senderAddress, emails, subject, msg); err == nil {
				resp.Error = 0
				resp.ErrorMsg = "Message" + msg + "sent via" + channel + "to" + emails
				if err := ret.Encode(resp); err != nil {
					log.Error(err.Error())
				}
			} else {
				resp.Error = 1
                                resp.ErrorMsg = err.Error()
                                if err := ret.Encode(resp); err != nil {
                                        log.Error(err.Error())
                                }	
			}
		} else {
			resp.Error = 1
			resp.ErrorMsg = "Failed to send message via" + channel
			if err := ret.Encode(resp); err != nil {
				log.Error(err.Error())
			}
		}
	default:
		resp.Error = 1
		resp.ErrorMsg = "No valid channel found"
		if err := ret.Encode(resp); err != nil {
			log.Error(err.Error())
		}
	}
}
