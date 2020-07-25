package login

import (
	"fmt"
	"log"
	"net/http"
)

func HandleRequest(r http.ResponseWriter) {
	log.Println("Login test")
	fmt.Fprintln(r, "Login test")
}
