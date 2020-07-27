package controler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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
	if r.Method != "POST" { // unit test needed
		http.Error(responseWriter, http.StatusText(405), 405)
		return
	}

	var loginReq loginReq
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(responseWriter, http.StatusText(400), 400)
		return
	}

	user, err := env.store.GetUserByUsername(strings.ToLower(loginReq.Username))
	if err != nil {
		http.Error(responseWriter, "User not found", 404)
		return
	}

	fmt.Fprintln(responseWriter, user)
}
