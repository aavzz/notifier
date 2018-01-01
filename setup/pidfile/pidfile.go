/*
Package pidfile checks if the process is already running
and manages PID files.
*/
package pidfile

import (
	. "github.com/aavzz/notifier/setup/cmdlnopts"
	. "github.com/aavzz/notifier/setup/syslog"
	"github.com/tabalt/pidfile"
	"os"
)

var p *pidfile.PidFile

// WritePid checks if the process is already running
// and tries to write PID file
// It runs in child only (stdout is unavailavle,
// logging has been already initialized)
func WritePid() {
	p = pidfile.NewPidFile(ConfigPidFile())
	oldpid, err := p.ReadPidFromFile(p.File)
	if err == nil && oldpid.ProcessExist() {
		SysLog.Err("Another process is already running")
		os.Exit(1)
	}

	if err := p.WritePidToFile(p.File, p.Pid); err != nil {
		SysLog.Err(err.Error())
		os.Exit(1)
	}
}

// RemovePidFile removes PID file.
func RemovePidFile() {
	p.Clear()
}
