package cmd

import (
	"github.com/aavzz/misc/pipe"
	"github.com/aavzz/notifier"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var smsc = &cobra.Command{
	Use:   "smsc",
	Short: "Sends an SMS via smsc",
	Long:  `Instructs notifyd to send SMS via smsc`,
	Run:   smscCommand,
}

func smscCommand(cmd *cobra.Command, args []string) {

	//read message from stdin (pipe)
	message, err := pipe.Read(800)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = notifier.NotifySMS(viper.GetString("smsc.url"), "smsc", viper.GetString("smsc.recipients"), message)
	if err != nil {
		log.Fatal(err.Error())
	}
}
