/*
Package notifier provides GO API to notifyd
*/
package notifier

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/url"
	"errors"
	"io/ioutil"
	"github.com/aavzz/daemon/log"
)

//Must be exportable (used for notifier response)
type NotifierResponse struct {
	Error    int
	ErrorMsg string
}

//NotifySMS sends an SMS via notifyd
//notifyd - notifyd https URL (e.g. "https://notifyd.somewhere.com/api1")
//channel - name of the SMS gateway (e.d. beeline)
//phones - comma-separatel list of cell-phones in international format (e.g. +71231234567,+71231234568)
//message - message to send
func NotifySMS(notifyd, channel, phones, message string) error {

	params := url.Values{
		"channel":    {channel},
		"recipients": {phones},
		"message":    {message},
	}

	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	resp, err := c.PostForm(notifyd, params)
	if err != nil {
		return err
	}
	if resp != nil {
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error(err.Error())
			}

			var v NotifierResponse
			if err := json.Unmarshal(body, &v); err != nil {
				return err
			}

			if v.Error != 0 {
				return errors.New(v.ErrorMsg)
			}
		} else {
			return errors.New(resp.Status)
		}
	} else {
		return errors.New("No response from notifyd")
	}
	return nil
}

//NotifyTelegram sends a message via a notifyd telegram bot
//notifyd - notifyd https URL (e.g. "https://notifyd.somewhere.com/api1")
//group - name os the group to send message to (as in the notifyd configuration file)
//message - message to send
func NotifyTelegram(notifyd, group, message string) error {

	params := url.Values{
		"channel":    {"telegram"},
		"recipients": {group},
		"message":    {message},
	}

	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	resp, err := c.PostForm(notifyd, params)
	if err != nil {
		return err
	}
	if resp != nil {
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error(err.Error())
			}

			var v NotifierResponse
			if err := json.Unmarshal(body, &v); err != nil {
				return err
			}

			if v.Error != 0 {
				return errors.New(v.ErrorMsg)
			}
		} else {
			return errors.New(resp.Status)
		}
	} else {
		return errors.New("No response from notifyd")
	}
	return nil
}

//NotifyEmail sends a message via a notifyd email facility
//notifyd - notifyd https URL (e.g. "https://notifyd.somewhere.com/api1")
//recipients - comma-separated list of email addresses
//subject - message subject
//sender_name - message sender name
//sender_addr - message sendet email address
//message - message to send
func NotifyEmail(notifyd, recipients, subject, sender_name, sender_address, message string) error {

	params := url.Values{
		"channel":        {"email"},
		"recipients":     {recipients},
		"subject":        {subject},
		"sender_name":    {sender_name},
		"sender_address": {sender_addr},
		"message":        {message},
	}

	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}

	resp, err := c.PostForm(notifyd, params)
	if err != nil {
		return err
	}
	if resp != nil {
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error(err.Error())
			}

			var v NotifierResponse
			if err := json.Unmarshal(body, &v); err != nil {
				return err
			}

			if v.Error != 0 {
				return errors.New(v.ErrorMsg)
			}
		} else {
			return errors.New(resp.Status)
		}
	} else {
		return errors.New("No response from notifyd")
	}
	return nil
}
