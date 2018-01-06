/*
notifyd sends messages via vifferent providers
(sms, email, etc.)
*/
package main

import (
	. "github.com/aavzz/notifier/setup"
	. "github.com/aavzz/notifier/setup/pidfile"
)

func main() {
	Setup()
	RemovePidFile()
}
