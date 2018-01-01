/*
Package setup initializes the process to run in the
background.
*/
package setup

import (
	"fmt"
	. "github.com/aavzz/notifier/rest"
	. "github.com/aavzz/notifier/setup/cfgfile"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
	. "github.com/aavzz/notifier/setup/daemon"
	. "github.com/aavzz/notifier/setup/pidfile"
	. "github.com/aavzz/notifier/setup/signal"
	. "github.com/aavzz/notifier/setup/syslog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Setup spawns child process, checks that everything is ok,
// does the necessary initialization and stops the parent process.
func Setup() {

	//create child process
	p, err := Daemonize()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//only entered in parent
	if p != nil {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM)

		//checks command line args, config file and double invocation
		//writes errors to stdout
		//and gets the coffin ready
		ParseCmdLine()
		InitLogging()

		<-sigterm
		os.Exit(0)
	}

	//parent never gets here

	//give the parent time to install signal handler
	//so we don't kill it prematurely
	time.Sleep(100 * time.Millisecond)

	//say good bye to parent
	ppid := os.Getppid()
	//we don't want to kill init
	if ppid > 1 {
		syscall.Kill(ppid, syscall.SIGTERM)
	}

	//child's output goes to /dev/null
	//we processed this in parent just to check for correctness
	//real configuration happens here
	ParseCmdLine()
	InitLogging()

	//rest of initialization
	WritePid()
	ReadConfig()
	SignalHandling()

	SysLog.Info("Server process started")

	InitHttp()
}
