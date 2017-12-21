package setup

import (
	"os"
	"github.com/gorilla/mux"
	"github.com/aavzz/rest/api1"
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
