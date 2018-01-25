package channels

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

// InitTelegram initializes telegram bot
func InitTelegram() {
	err error
	bot, err := tgbotapi.NewBotAPI(viper.GetString("telegram.token"))
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
