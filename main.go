package main

/*
 * stdout is unavailable after Setup()
 */

import (
	"os"
	"net/http"
	. "github.com/aavzz/notifier/setup/setup"
)

func main() {

	Setup()
	defer RemovePidFile()
	SysLog.Info("Successfull start")

}
