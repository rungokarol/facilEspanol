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
	if r.Method != "POST" { //unit test needed
		http.Error(responseWriter, http.StatusText(405), 405)
		return
	}

	var loginReq loginReq
	json.NewDecoder(r.Body).Decode(&loginReq) //handle errors

	user, err := env.store.GetUserByUsername(strings.ToLower(loginReq.Username))
	if err != nil {
		log.Println(err)
	} else {
		log.Println(user)
	}

	fmt.Fprintf(responseWriter, "user login")
}
