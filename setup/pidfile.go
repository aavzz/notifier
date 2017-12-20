package setup

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"fmt"
	"github.com/tabalt/pidfile"
	"os"
)

var p *pidfile.PidFile

func writePid() {
	p = pidfile.NewPidFile(cmdLnOpts.pidfile)
	oldpid, err := p.ReadPidFromFile(p.File)
	if err == nil && oldpid.ProcessExist() {
		fmt.Println("Another process is already running")
		os.Exit(1)
	}

	//avoid creating pidfile in parent
	daemonState := os.Getenv("_GO_DAEMON_STATE")
	if daemonState == "" {
		if err := p.WritePidToFile(p.File, p.Pid); err != nil {
			SysLog.Err("Cannot create pidfile")
			os.Exit(1)
		}
	}
}

func RemovePidFile() {
	p.Clear()
}
