package setup

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/aavzz/notifier/rest/api1"
)

func initHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/api1", api1.handler).Methods("GET")
  
	err := http.ListenAndServe(ConfigAddress(), r)
	if err != nil {
		SysLog.Err(err.Error())
		os.Exit(1)
	}
  
}
