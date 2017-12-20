package setup

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"syscall"
)

func daemonize() (*os.Process, error) {

	daemonState := os.Getenv("_GO_DAEMON_STATE")
	switch daemonState {
	case "":
		syscall.Umask(0022)
		syscall.Setsid()
		os.Setenv("_GO_DAEMON_STATE", "1")
	case "1":
		os.Setenv("_GO_DAEMON_STATE", "")
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
