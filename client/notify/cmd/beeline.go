package cmd

import (
	"github.com/aavzz/misc/pipe"
	"github.com/aavzz/notifier"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

var beeline = &cobra.Command{
	Use:   "beeline",
	Short: "Sends an SMS via beeline",
	Long:  `Instructs notifyd to send SMS via beeline`,
	Run:   beelineCommand,
}

func beelineCommand(cmd *cobra.Command, args []string) {

	//read message from stdin (pipe)
	message, err := pipe.Read(480)
	if err != nil {
		log.Fatal(err.Error())
	}

	err := notifier.NotifySMS(viper.GetString("beeline.url"), "beeline", viper.GetString("beeline.recipients"), message)
	if err != nil {
		log.Fatal(err.Error())
	}
}
