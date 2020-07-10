package main

import (
  "fmt"
  "log"
  "net/http"
)

func handler(responseWriter http.ResponseWriter, r *http.Request) {
    log.Println("request received")

    fmt.Fprintf(responseWriter, "Hello %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

