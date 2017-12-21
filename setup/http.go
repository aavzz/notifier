package setup

import (
  "github.com/gorilla/mux"
)

func initHttp() {
  r := mux.NewRouter()
  r.HandleFunc("/api-1", handler).Methods("GET")
  
  err := http.ListenAndServe(ConfigAddress(), r)
	if err != nil {
		SysLog.Err(err.Error())
		os.Exit(1)
	}
  
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
