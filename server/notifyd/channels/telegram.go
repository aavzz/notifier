package channels

import (
	"github.com/spf13/viper"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/aavzz/daemon/log"
	"net/http"
	"crypto/tls"
)

var bot *tgbotapi.BotAPI

// InitTelegram initializes telegram bot
func InitTelegram() {
	var err error
	c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	bot, err := tgbotapi.NewBotAPIWithClient(viper.GetString("telegram.token"), c)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("Telegram bot authorized on account " + bot.Self.UserName)
		go respond()
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
	}
	return nil
}

///////////////////////////////////////////////////////////////

func respond() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Error(err.Error())
	}

	for {
		for update := range updates {
			if update.Message == nil {
				continue
			}
			SendMessageTelegram(update.Message.Chat.ID, "Тишина в библиотеке!!!")
		}
	}
}
