/*
Package cmd implements notify commands and flags
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var notify = &cobra.Command{
	Use:   "notify",
	Short: "notify is a client for notifyd",
	Long:  `A web-application client for notifyd`,
	Run:   notifyCommand,
}

func notifyCommand(cmd *cobra.Command, args []string) {
	//nothing here
}

// Execute starts notify execution
func Execute() {
        beeline.Flags().StringP("recipients", "r", "", "recipient list")
        beeline.Flags().StringP("message", "m", "", "message to send")
        beeline.Flags().StringP("url", "u", "", "url to query")
	viper.BindPFlag("beeline.recipients", beeline.Flags().Lookup("recipients"))
	viper.BindPFlag("beeline.message", beeline.Flags().Lookup("message"))
	viper.BindPFlag("beeline.url", beeline.Flags().Lookup("url"))

        email.Flags().StringP("recipients", "r", "", "recipient list")
        email.Flags().StringP("message", "m", "", "message to send")
        email.Flags().StringP("sender-name", "n", "", "sender name")
        email.Flags().StringP("sender-address", "a", "", "sender address")
        email.Flags().StringP("subject", "s", "", "email subject")
        email.Flags().StringP("url", "u", "", "url to query")
	viper.BindPFlag("email.recipients", email.Flags().Lookup("recipients"))
	viper.BindPFlag("email.message", email.Flags().Lookup("message"))
	viper.BindPFlag("email.sender-name", email.Flags().Lookup("sender-name"))
	viper.BindPFlag("email.sender-address", email.Flags().Lookup("sender-address"))
	viper.BindPFlag("email.subject", email.Flags().Lookup("subject"))
	viper.BindPFlag("email.url", email.Flags().Lookup("url"))

	notify.AddCommand(beeline)
	notify.AddCommand(email)

	if err := notify.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
