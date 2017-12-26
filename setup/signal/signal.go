package signal

import (
	"os"
	"syscall"
	"os/signal"
        . "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/pidfile"
	. "github.com/aavzz/notifier/setup/cfgfile"
)

func SignalHandling() {

	sighup := make(chan os.Signal, 1)
	sigint := make(chan os.Signal, 1)
	sigquit := make(chan os.Signal, 1)
	sigterm := make(chan os.Signal, 1)

	signal.Notify(sighup, syscall.SIGHUP)
	signal.Notify(sigint, syscall.SIGINT)
	signal.Notify(sigquit, syscall.SIGQUIT)
	signal.Notify(sigterm, syscall.SIGTERM)

	go func() {
		for {
			<-sighup
			SysLog.Info("SIGHUP received, re-reading configuration file")
			ReadConfig()
		}
	}()

	go func() {
		<-sigint
		SysLog.Info("SIGINT received, exitting")
		RemovePidFile()
		os.Exit(0)
	}()

	go func() {
		<-sigquit
		SysLog.Info("SIGQUIT received, exitting")
		RemovePidFile()
		os.Exit(0)
	}()

	go func() {
		<-sigterm
		SysLog.Info("SIGTERM received, exitting")
		RemovePidFile()
		os.Exit(0)
	}()

}
