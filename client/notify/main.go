package main

import (
	"github.com/aavzz/notifier/client/notify/cmd"
	"os"
	"log"
)

func main() {
	//check if we are invoked as a pipe
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err.Error())	
	}
	if fi.Mode() & os.ModeNamedPipe == 0 {
		log.Fatal("stdin not connected to pipe")
	}
	
	cmd.Execute()
}
