package channels

import (
	"gopkg.in/gomail.v2"
	"regexp"
)

// SendMessageEmail sends message via local email server
func SendMessageEmail(senderName, senderAddress, recipients, subject, message string) error {

	re := regexp.MustCompile(`\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`)
	emails := re.FindAllString(recipients, -1)
	senderAddr := re.FindAllString(senderAddress, 1)
	l := len(message)
	if l > 10000 {
		l = 10000
	}
	message = message[:l]

	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(senderAddr[0], senderName)},
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
