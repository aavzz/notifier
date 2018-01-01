/*
Package cmdlnopts parses command line options and
stores information for later use.
*/
package cmdlnopts

import (
	"github.com/pborman/getopt/v2"
)

// commandLineOptions is where all the information
// obtained from the command line is stored.
type commandLineOptions struct {
	address string
	cfgfile string
	pidfile string
}

var cmdLnOpts commandLineOptions

// ParseCmdLine parses command line options and fills the internal structure.
// In case of an error the process is stoped.
// It runs both in parent and child. Parent's output goes to stdout,
// child's to /dev/null
func ParseCmdLine() {
	a := getopt.StringLong("addr", 'a', "127.0.0.1:8080", "address and port to bind to")
	c := getopt.StringLong("cfg", 'c', "/etc/notifyd.conf", "configuration file")
	p := getopt.StringLong("pid", 'p', "/var/run/notifyd.pid", "PID file")
	getopt.Parse()
	cmdLnOpts.address = *a
	cmdLnOpts.cfgfile = *c
	cmdLnOpts.pidfile = *p
}

//ConfigPidFile returns the name of PID file
func ConfigAddress() string {
	return cmdLnOpts.address
}

//ConfigCFGFIle returns the name of PID file
func ConfigPidFile() string {
	return cmdLnOpts.pidfile
}

//ConfigCFGFIle returns the name of configuration file
func ConfigCfgFile() string {
	return cmdLnOpts.cfgfile
}
