package main

/*
 * stdout is unavailable after Setup()
 */

import (
	"os"
	"net/http"
	. "github.com/aavzz/notifier/setup"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/pidfile"
)

func main() {

	Setup()
	defer RemovePidFile()
	SysLog.Info("Successfull start")

}
