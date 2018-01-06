/*
Package rest implements REST interface of notifyd.
*/
package rest

import (
	"context"
	"github.com/aavzz/daemon/log"
	"github.com/aavzz/daemon/pid"
	"github.com/aavzz/daemon/signal"
	"github.com/aavzz/notifier/server/notifyd/rest/api1"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

// InitHttpp sets up router.
func InitHttp() {
	r := mux.NewRouter()
	r.HandleFunc("/api1", api1.Handler).Methods("POST")

	s := &http.Server{
		Addr:     viper.GetString("address"),
		Handler:  r,
		ErrorLog: log.Logger("notifyd"),
	}

	if viper.GetBool("daemonize") == true {
		signal.Quit(func() {
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			log.Info("SIGQUIT received, exitting gracefully")
			s.Shutdown(ctx)
			pid.Remove()
		})
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
