package main

/*
 * stdout is unavailable after Setup()
 */

import (
	"net/http"
	. "github.com/aavzz/notifier/setup"
)

func main() {

	Setup()
	defer RemovePidFile()
	SysLog.Info("starting daemon...")

	//just some event loop
	//newer mind that it does not do anything useful 
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(ConfigAddress(), nil)
	if err != nil {
		SysLog.Err(err.Error())
		os.Exit(1)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	SysLog.Info("Hello World!")
}
