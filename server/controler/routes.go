package controler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type loginReq struct {
	Username string
	Password string
}

func (env *Env) DefaultRoot(responseWriter http.ResponseWriter, r *http.Request) {
	log.Println("request received")
	fmt.Fprintf(responseWriter, "Hello %s!", r.URL.Path[1:])
}

func (env *Env) Login(responseWriter http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(responseWriter, http.StatusText(405), 405)
		return
	}

	var loginReq loginReq
	json.NewDecoder(r.Body).Decode(&loginReq)

	fmt.Fprintf(responseWriter, "user login")
}
