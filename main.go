package main

/*
 * stdout is unavailable after Setup()
 */

import (
	. "github.com/aavzz/notifier/setup"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/pidfile"
)

func main() {
	Setup()
	RemovePidFile()
}
