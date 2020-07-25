package login

import (
	"fmt"
	"log"
	"net/http"
)

func LoginHandler(responseWriter http.ResponseWriter, r *http.Request) {
	log.Println("Login test")
	fmt.Fprintln(responseWriter, "Login test")
}
