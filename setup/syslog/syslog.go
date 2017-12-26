package syslog

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"fmt"
	"log/syslog"
)

var SysLog *syslog.Writer

func InitLogging() {
	var err error
	SysLog, err = syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "notifier")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

