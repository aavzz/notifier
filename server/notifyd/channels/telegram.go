package channels

import (
	"crypto/tls"
	"errors"
	"github.com/aavzz/daemon/log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"net/http"
)

var bot *tgbotapi.BotAPI

// InitTelegram initializes telegram bot
func InitTelegram() {
	var err error
	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	bot, err = tgbotapi.NewBotAPIWithClient(viper.GetString("telegram.token"), c)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("Telegram bot authorized on account " + bot.Self.UserName)
	}
}

///////////////////////////////////////////////////////////////

// SendMessageTelegram sends message to the specified telegram group
func SendMessageTelegram(chatID int64, message string) error {

	if bot != nil {
		_, err := bot.Send(tgbotapi.NewMessage(chatID, message))
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Bot is not connected")
}
