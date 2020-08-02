package controler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
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
		http.Error(responseWriter, http.StatusText(500), 500)
		return
	} else if user == nil {
		http.Error(responseWriter, "User not found", 404)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password)) ; err != nil {
		http.Error(responseWriter, "Wrong username or password", 403)
		return
	}

	// 1. Zahaszowac haslo z requsta 	DONE
	// 2. Porownac z hashem w rekordzie DONE
	// 3. Stworz i wyslij JWT			TODO



	fmt.Fprintln(responseWriter, user)
}
