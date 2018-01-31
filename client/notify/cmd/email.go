package cmd

import (
	"github.com/aavzz/misc/pipe"
	"github.com/aavzz/notifier"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var email = &cobra.Command{
	Use:   "email",
	Short: "Sends an email",
	Long:  `Instructs notifyd to send message via local mailserver`,
	Run:   emailCommand,
}

func emailCommand(cmd *cobra.Command, args []string) {

	//read message from stdin (pipe)
	message, err := pipe.Read(1024)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = notifier.NotifyEmail(viper.GetString("email.url"), viper.GetString("email.recipients"),
		viper.GetString("email.subject"), viper.GetString("email.sender-name"),
		viper.GetString("email.sender-address"), message)
	if err != nil {
		log.Fatal(err.Error())
	}
}
