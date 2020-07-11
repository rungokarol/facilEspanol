package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/rungokarol/facilEspanol/src/login"
)

func handler(responseWriter http.ResponseWriter, r *http.Request) {
    log.Println("request received")

    fmt.Fprintf(responseWriter, "Hello %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/login", login.LoginHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

