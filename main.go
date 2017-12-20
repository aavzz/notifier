package main

/*
 * stdout is unavailable after Setup()
 */

import (
	"net/http"
	. "github.com/aavzz/stub-server/setup"
)

func main() {

	Setup()
	defer RemovePidFile()
	SysLog.Info("starting daemon...")

	//just some event loop
	//newer mind that it does not do anything useful 
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	SysLog.Info("Hello World!")
}
