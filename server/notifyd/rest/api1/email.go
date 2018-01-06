package api1

import (
	"gopkg.in/gomail.v2"
)

func sendMessageEmail(senderName string, senderAddress string, emails []string, subject string, message string) error {

	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(senderAddress, senderName)},
		"To":      emails,
		"Subject": {subject},
	})
	m.SetBody("text/plain", message)

	d := gomail.Dialer{Host: "localhost", Port: 25}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
