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
	beeline.Flags().StringP("recipients", "r", "", "comma-delimited phone list in international format")
	beeline.Flags().StringP("url", "u", "", "notifyd URL to POST the message to")
	viper.BindPFlag("beeline.recipients", beeline.Flags().Lookup("recipients"))
	viper.BindPFlag("beeline.url", beeline.Flags().Lookup("url"))

	smsc.Flags().StringP("recipients", "r", "", "comma-delimited phone list in international format")
	smsc.Flags().StringP("url", "u", "", "notifyd URL to POST the message to")
	viper.BindPFlag("smsc.recipients", smsc.Flags().Lookup("recipients"))
	viper.BindPFlag("smsc.url", smsc.Flags().Lookup("url"))

	websms.Flags().StringP("recipients", "r", "", "comma-delimited phone list in international format")
	websms.Flags().StringP("url", "u", "", "notifyd URL to POST the message to")
	viper.BindPFlag("websms.recipients", websms.Flags().Lookup("recipients"))
	viper.BindPFlag("websms.url", websms.Flags().Lookup("url"))

	telegram.Flags().StringP("group", "r", "", "telegram group name in the configuration file")
	telegram.Flags().StringP("url", "u", "", "notifyd URL to POST the message to")
	viper.BindPFlag("telegram.group", telegram.Flags().Lookup("group"))
	viper.BindPFlag("telegram.url", telegram.Flags().Lookup("url"))

	email.Flags().StringP("recipients", "r", "", "comma-delimited email address list")
	email.Flags().StringP("sender-name", "n", "", "sender name")
	email.Flags().StringP("sender-address", "a", "", "sender address")
	email.Flags().StringP("subject", "s", "", "email subject")
	email.Flags().StringP("url", "u", "", "notifyd URL to POST the message to")
	viper.BindPFlag("email.recipients", email.Flags().Lookup("recipients"))
	viper.BindPFlag("email.sender-name", email.Flags().Lookup("sender-name"))
	viper.BindPFlag("email.sender-address", email.Flags().Lookup("sender-address"))
	viper.BindPFlag("email.subject", email.Flags().Lookup("subject"))
	viper.BindPFlag("email.url", email.Flags().Lookup("url"))

	notify.AddCommand(beeline)
	notify.AddCommand(smsc)
	notify.AddCommand(websms)
	notify.AddCommand(telegram)
	notify.AddCommand(email)

	if err := notify.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
