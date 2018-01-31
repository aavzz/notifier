package cmd

import (
	"github.com/aavzz/misc/pipe"
	"github.com/aavzz/notifier"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var telegram = &cobra.Command{
	Use:   "telegram",
	Short: "Sends a message to telegram group",
	Long:  `Instructs notifyd to send message to telegram group`,
	Run:   telegramCommand,
}

func telegramCommand(cmd *cobra.Command, args []string) {

	//read message from stdin (pipe)
	message, err := pipe.Read(800)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = notifier.NotifyTelegram(viper.GetString("telegram.url"), viper.GetString("telegram.group"), message)
	if err != nil {
		log.Fatal(err.Error())
	}
}
