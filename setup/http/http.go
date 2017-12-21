package http

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/aavzz/notifier/rest/api1"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/cmdlnopts"
)

func InitHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/api1", api1.Handler).Methods("GET")
  
	err := http.ListenAndServe(ConfigAddress(), r)
	if err != nil {
		SysLog.Err(err.Error())
		os.Exit(1)
	}
}

