package controler

import (
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type loginReq struct {
	Username string
	Password string
}

type loginResp struct {
	Token string `json:"token"`
}

var minLength = 3

func (env *Env) Login(responseWriter http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var loginReq loginReq
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := env.store.GetUserByUsername(strings.ToLower(loginReq.Username))
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	} else if user == nil {
		http.Error(responseWriter, "User not found", http.StatusNotFound) //not sure if correct status
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash),
		[]byte(loginReq.Password)); err != nil {
		http.Error(responseWriter,
			"Wrong username or password",
			http.StatusForbidden)
		return
	}

	token, err := createJwt(user.Username)
	if err != nil {
		http.Error(responseWriter,
			"Error creating JWT",
			http.StatusInternalServerError)
		return
	}

	responseJson, err := json.Marshal(loginResp{Token: token})
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	//status 201 - created; 202- accepted
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(responseJson)
}
