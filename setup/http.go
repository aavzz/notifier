package setup

import (
  "github.com/gorilla/mux"
)

func initHttp() {
  r := mux.NewRouter()
  r.HandleFunc("/api-1", handler).Methods("GET")
  http.ListenAndServe(":8080", r)
}
