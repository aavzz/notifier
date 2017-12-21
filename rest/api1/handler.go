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
			re := regexp.MustCompile(`+7\d{10}`)
			phones := re.FindAllString(recipients, 5)
			l := len(message)
			if l > 480 {
				l = 480
			}
			msg := message[:l]
			if {
			
			}
		case "email":
		
		case "telegram":

		default:
		
	}
	
	
	
	fmt.Fprintf(w, "Hello World!")
}
