/*
notifyd sends messages via different providers
(sms, email, etc.)
*/
package main

import (
	"github.com/aavzz/notifier/server/notifyd/cmd"
)

func main() {
	cmd.Execute()
}
