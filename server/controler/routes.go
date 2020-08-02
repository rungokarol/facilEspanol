package controler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rungokarol/facilEspanol/model"
	"golang.org/x/crypto/bcrypt"
)

type loginReq struct {
	Username string
	Password string
}

type loginResponse struct {
	token string
}

type registerReq struct {
	Username string
	Password string
}

var minLength = 3

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

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginReq.Password)); err != nil {
		http.Error(responseWriter, "Wrong username or password", 403)
		return
	}

	token, err := createJwt(user.Username)
	if err != nil {
		http.Error(responseWriter, "Error creating JWT", 500)
		return
	}

	response := loginResponse{
		token: token,
	}

	// TODO: check if it's converting to JSON
	fmt.Fprintln(responseWriter, response)
}

func (env *Env) Register(responseWriter http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // unit test needed
		http.Error(responseWriter, http.StatusText(405), 405)
		return
	}

	var registerReq registerReq
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		http.Error(responseWriter, http.StatusText(400), 400)
		return
	}

	if len(registerReq.Username) < minLength || len(registerReq.Password) < minLength {
		http.Error(responseWriter, "Username or password too short", 400)
		return
	}

	isPresent, err := env.store.IsUserPresent(strings.ToLower(registerReq.Username))
	if err != nil {
		http.Error(responseWriter, http.StatusText(500), 500)
		return
	} else if isPresent {
		http.Error(responseWriter, "User with given username already exists", 400)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(responseWriter, http.StatusText(500), 500)
		return
	}

	model := model.User{
		Username:     strings.ToLower(registerReq.Username),
		PasswordHash: string(hashedPassword),
	}

	if err := env.store.CreateUser(&model); err != nil {
		http.Error(responseWriter, http.StatusText(500), 500)
		return
	}
}
