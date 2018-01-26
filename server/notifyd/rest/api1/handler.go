/*
Package api1 implements version 1 of notifyd API.
*/
package api1

import (
	"encoding/json"
	"github.com/aavzz/daemon/log"
	"github.com/aavzz/notifier/server/notifyd/channels"
	"github.com/spf13/viper"
	"net/http"
)

// JResponse holds notifyd response
// Must be exportable
type JResponse struct {
	Error    int
	ErrorMsg string
}

// Handler calls the right function to send message via specified channel.
func Handler(w http.ResponseWriter, r *http.Request) {

	channel := r.FormValue("channel")
	message := r.FormValue("message")
	recipients := r.FormValue("recipients")

	switch channel {
	case "beeline":
		if recipients != "" && message != "" {
			if err := channels.SendMessageBeeline(viper.GetString("beeline.login"),
				viper.GetString("beeline.password"),
				viper.GetString("beeline.sender"),
				recipients, message); err == nil {
				reportSuccess(w, message, channel, recipients)
			} else {
				reportError(w, err)
			}
		} else {
			reportErrorString(w, "Failed to send message via "+channel)
		}
	case "smsc":
		if recipients != "" && message != "" {
			if err := channels.SendMessageSmsc(viper.GetString("smsc.login"),
				viper.GetString("smsc.password"),
				viper.GetString("smsc.sender"),
				recipients, message); err == nil {
				reportSuccess(w, message, channel, recipients)
			} else {
				reportError(w, err)
			}
		} else {
			reportErrorString(w, "Failed to send message via "+channel)
		}
	case "websms":
		if recipients != "" && message != "" {
			if err := channels.SendMessageWebsms(viper.GetString("websms.login"),
				viper.GetString("websms.password"),
				viper.GetString("websms.sender"),
				recipients, msg); err == nil {
				reportSuccess(w, message, channel, recipients)
			} else {
				reportError(w, err)
			}
		} else {
			reportErrorString(w, "Failed to send message via "+channel)
		}
	case "telegram":
		if recipients != "" && message != "" {
			if err := channels.SendMessageTelegram(viper.GetInt64("telegram."+recipients+"_chatid"), message); err == nil {
				reportSuccess(w, message, channel, recipients)
			} else {
				reportError(w, err)
			}
		} else {
			reportErrorString(w, "Failed to send message via "+channel)
		}
	case "email":
		senderName := r.FormValue("sender_name")
		senderAddr := r.FormValue("sender_address")
		subject := r.FormValue("subject")

		if recipients != "" && message != "" {
			if err := channels.SendMessageEmail(senderName, senderAddr, recipients, subject, message); err == nil {
				reportSuccess(w, message, channel, recipients)
			} else {
				reportError(w, err)
			}
		} else {
			reportErrorString(w, "Failed to send message via "+channel)
		}
	default:
		reportErrorString(w, "No valid channel found")
	}
}

//////////////////////////////////////////////////////////////////////////////

func reportError(w http.ResponseWriter, e error) {
	ret := json.NewEncoder(w)
	var resp JResponse
	resp.Error = 1
	resp.ErrorMsg = e.Error()
	if err := ret.Encode(resp); err != nil {
		log.Error(err.Error())
	}
}

//////////////////////////////////////////////////////////////////////////////

func reportErrorString(w http.ResponseWriter, e string) {
	ret := json.NewEncoder(w)
	var resp JResponse
	resp.Error = 1
	resp.ErrorMsg = e
	if err := ret.Encode(resp); err != nil {
		log.Error(err.Error())
	}
}

//////////////////////////////////////////////////////////////////////////////

func reportSuccess(w http.ResponseWriter, msg, channel, recipients string) {
	ret := json.NewEncoder(w)
	var resp JResponse
	resp.Error = 0
	resp.ErrorMsg = "Message `" + msg + "` sent via " + channel + " to " + recipients
	if err := ret.Encode(resp); err != nil {
		log.Error(err.Error())
	}
}
