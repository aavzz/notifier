package api1

include (
  "fmt"
  "github.com/gorilla/http"
)

func sendMessage(numbers string, message string) {
  c := new(http.Client) 

  status, _, r, err := c.Post("http://www.example.com", nil, "qwertyu")
  if err != nil {
    SysLog.Err(err.Error())
  }
  if r != nil {
    defer r.Close()
  }
  SysLog.Info(fmt.Sprintf("Post result: %v", status))
  
}
