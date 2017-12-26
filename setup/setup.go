package setup

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"os"
	"fmt"
	"time"
	"syscall"
	"os/signal"
	. "github.com/aavzz/notifier/setup/http"
	. "github.com/aavzz/notifier/setup/daemon"
	. "github.com/aavzz/notifier/setup/signal"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/pidfile"
	. "github.com/aavzz/notifier/setup/cfgfile"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
)

func Setup() {

	//create child process
	p, err := Daemonize()
	if err != nil {
		fmt.Printf("Cannot daemonize: %s/n", err)
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
		WritePid()
		ReadConfig()
		
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
	WritePid()
	ReadConfig()
	
	SysLog.Info("Server process started")
	
	//rest of initialization
	SignalHandling()
	InitHttp()
}
