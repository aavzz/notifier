package pidfile

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"fmt"
	"github.com/tabalt/pidfile"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
)

var p *pidfile.PidFile

func WritePid() {
	p = pidfile.NewPidFile(ConfigPidFile())
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

