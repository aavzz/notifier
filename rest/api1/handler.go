package api1

import (
	"fmt"
	"regexp"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	channel := r.FormValue("channel")
	recipients := r.FormValue("recipients")
	message := r.FormValue("message")

	switch channel {
		case "beeline":
			re := regexp.MustCompile(`+7\d{10},`)
		//check recipients
		//check message (480 symbols)
		case "email":
		
		case "telegram":

		default:
		
	}
	
	
	
	fmt.Fprintf(w, "Hello World!")
}
