package setup

/*
 * this code runs both in parent and child
 * so beware of stdout availability (parent only)
 */

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Setup() {

	//create child process
	p, err := daemonize()
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
		parseCmdLine()
		writePid()
		readConfig()
		initLogging()

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
	parseCmdLine()
	writePid()
	readConfig()
	initLogging()

	//final touch
	signalHandling()

}

func ConfigPidFile() string {
	return cmdLnOpts.pidfile
}

func ConfigCfgFile() string {
	return cmdLnOpts.cfgfile
}

func ConfigSmtpServerAddress() string {
	return cfgFile.smtp.address
}

func ConfigSmtpServerPort() uint16 {
	return cfgFile.smtp.port
}
