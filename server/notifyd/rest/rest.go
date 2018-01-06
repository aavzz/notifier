/*
Package rest implements REST interface of notifyd.
*/
package rest

import (
	"github.com/aavzz/notifier/rest/api1"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
	. "github.com/aavzz/notifier/setup/syslog"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

// InitHttpp sets up router.
func InitHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/api1", api1.Handler).Methods("GET")

	err := http.ListenAndServe(ConfigAddress(), r)
	if err != nil {
		SysLog.Err(err.Error())
		os.Exit(1)
	}
}
