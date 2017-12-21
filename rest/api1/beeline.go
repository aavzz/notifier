package api1

include (
  "fmt"
  "net/url"
  "github.com/gorilla/http"
)

func sendMessage(numbers string, message string) {
    
  parameters := url.Values{}
  parameters.Add("user", "")
  parameters.Add("pass", "")
  parameters.Add("action", "post_sms")
  parameters.Add("target", numbers)
  parameters.Add("message", message)
  
  //Content-Type: application/x-www-form-urlencoded; charset=windows-1251
  
  c := new(http.Client)

  status, _, r, err := c.Post("https://beeline.amega-inform.ru/sendsms/", nil, parameters.Encode())
  if err != nil {
    SysLog.Err(err.Error())
  }
  if r != nil {
    defer r.Close()
  }
  SysLog.Info(fmt.Sprintf("Post result: %v", status))
  
}
