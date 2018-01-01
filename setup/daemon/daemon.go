/*
Package daemon sets up the process to run in the background.
Since fork() is unavailable in Go, child process is started
as a regular process. A special environment variable is used
to syncronize parent's and child's behaviour.
*/
package daemon

import (
	"os"
	"syscall"
)

// Daemonize starts child process.
func Daemonize() (*os.Process, error) {

	daemonState := os.Getenv("_NOTIFY_DAEMON_STATE")
	switch daemonState {
	case "":
		syscall.Umask(0022)
		syscall.Setsid()
		os.Setenv("_NOTIFY_DAEMON_STATE", "1")
	case "1":
		os.Setenv("_NOTIFY_DAEMON_STATE", "")
		return nil, nil
	}

	var attrs os.ProcAttr
	f, err := os.Open("/dev/null")
	if err != nil {
		return nil, err
	}
	attrs.Files = []*os.File{f, f, f}

	exec_path, err := os.Executable()
	if err != nil {
		return nil, err
	}

	p, err := os.StartProcess(exec_path, os.Args, &attrs)
	if err != nil {
		return nil, err
	}

	return p, nil
}
