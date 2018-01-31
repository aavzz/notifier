package cmd

import (
	"github.com/aavzz/misc/pipe"
	"github.com/aavzz/notifier"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

var websms = &cobra.Command{
	Use:   "websms",
	Short: "Sends an SMS via websms",
	Long:  `Instructs notifyd to send SMS via websms`,
	Run:   websmsCommand,
}

func websmsCommand(cmd *cobra.Command, args []string) {

	//read message from stdin (pipe)
	message, err := pipe.Read(800)
	if err != nil {
		log.Fatal(err.Error())
	}

	err := notifier.NotifySMS(viper.GetString("websms.url"), "websms", viper.GetString("websms.recipients"), message)
	if err != nil {
		log.Fatal(err.Error())
	}
}
