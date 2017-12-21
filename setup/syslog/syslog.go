package syslog

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"fmt"
	"log/syslog"
	"os"
)

var SysLog *syslog.Writer

func InitLogging() {
	var err error
	SysLog, err = syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "notifier")
	if err != nil {
		fmt.Printf("Cannot initialize logging: %s\n", err)
		os.Exit(1)
	}
}

