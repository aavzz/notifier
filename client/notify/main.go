package main

import (
	"github.com/aavzz/misc/pipe"
	"github.com/aavzz/notifier/client/notify/cmd"
	"log"
)

func main() {
	// check if stdin is connected to a pipe,
	ok, err := pipe.CheckStdin()
	if err != nil {
		log.Fatal(err.Error())
	}
	if ok != true {
		log.Fatal("stdin not connected to pipe")		
	}	
	
	cmd.Execute()
}
