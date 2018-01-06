/*
Package cmd implements notify commands and flags
*/
package cmd

import (
	"github.com/spf13/cobra"
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

        email.Flags().StringP("recipients", "r", "", "recipient list")
        email.Flags().StringP("message", "m", "", "message to send")
        email.Flags().StringP("sender-name", "n", "", "sender name")
        email.Flags().StringP("sender-address", "a", "", "sender address")
        email.Flags().StringP("subject", "s", "", "email subject")
        email.Flags().StringP("url", "u", "", "url to query")

	notify.AddCommand(beeline)
	notify.AddCommand(email)

	if err := notify.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
