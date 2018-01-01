/*
Package syslog sets up logging to local syslog
*/
package syslog

import (
	"fmt"
	"log/syslog"
	"os"
)

// SysLog methods do the logging
var SysLog *syslog.Writer

// InitLogging tries to get in touch with local syslog.
// It writes error message to stdout and stops the
// process in case of failure.
// It runs both in parent and child
// parent's output goes to stdout,
// child's to /dev/null
func InitLogging() {
	var err error
	SysLog, err = syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "notifyd")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
