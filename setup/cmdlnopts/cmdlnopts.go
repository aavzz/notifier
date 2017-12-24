package cmdlnopts

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"github.com/pborman/getopt/v2"
)

type commandLineOptions struct {
	address string
	cfgfile string
	pidfile string
}

var cmdLnOpts commandLineOptions

// Package exported objects

func ParseCmdLine() {
	a := getopt.StringLong("addr", 'a', "0.0.0.0:8080", "address and port to bind to")
	c := getopt.StringLong("cfg", 'c', "/etc/notifier.conf", "configuration file")
	p := getopt.StringLong("pid", 'p', "/var/run/notifier.pid", "PID file")
	getopt.Parse()
	cmdLnOpts.address = *a
	cmdLnOpts.cfgfile = *c
	cmdLnOpts.pidfile = *p
}

func ConfigAddress() string {
	return cmdLnOpts.address
}

func ConfigPidFile() string {
	return cmdLnOpts.pidfile
}

func ConfigCfgFile() string {
	return cmdLnOpts.cfgfile
}
