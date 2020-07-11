package login

import (
  "fmt"
  "net/http"
  "log"
)

func LoginHandler(responseWriter http.ResponseWriter, r *http.Request) {
  log.Println("Login test")

  fmt.Fprintln(responseWriter, "Login test")
}
