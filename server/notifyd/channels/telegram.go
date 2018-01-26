package channels

import (
	"github.com/spf13/viper"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/aavzz/daemon/log"
	"fmt"
	"net/http"
	"crypto/tls"
	"errors"
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
		log.Info(fmt.Sprintf("sending %s to %d", message, chatID))
		_, err := bot.Send(tgbotapi.NewMessage(chatID, message))
		if err != nil {
			return err
		}
	}
	return errors.New("Bot is not connected")
}
