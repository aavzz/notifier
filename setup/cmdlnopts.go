package setup

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

func parseCmdLine() {
	a := getopt.StringLong("addr", 'a', "127.0.0.1:8080", "address and port to bind to")
	c := getopt.StringLong("cfg", 'c', "/etc/notifier.conf", "configuration file")
	p := getopt.StringLong("pid", 'p', "/var/run/notifier.pid", "PID file")
	getopt.Parse()
	cmdLnOpts.address = *a
	cmdLnOpts.cfgfile = *c
	cmdLnOpts.pidfile = *p
}
