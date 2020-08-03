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
	Token string `json:"token"`
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
	if r.Method != "POST" {
		http.Error(responseWriter,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)
		return
	}

	var loginReq loginReq
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(responseWriter,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	user, err := env.store.GetUserByUsername(strings.ToLower(loginReq.Username))
	if err != nil {
		http.Error(responseWriter,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
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

	response := loginResponse{
		Token: token,
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(responseWriter,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	//status 201 - created; 202- accepted
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(responseJson)
}

func (env *Env) Register(responseWriter http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(responseWriter,
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)
		return
	}

	var registerReq registerReq
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		http.Error(responseWriter,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	if len(registerReq.Username) < minLength || len(registerReq.Password) < minLength {
		http.Error(responseWriter,
			"Username or password too short",
			http.StatusBadRequest)
		return
	}

	isPresent, err := env.store.IsUserPresent(strings.ToLower(registerReq.Username))
	if err != nil {
		http.Error(responseWriter,
			http.StatusText(http.StatusInternalServerError), //check if good status maybe unauthorized?
			http.StatusInternalServerError)
		return
	} else if isPresent {
		http.Error(responseWriter,
			"User with given username already exists",
			http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(responseWriter,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	model := model.User{
		Username:     strings.ToLower(registerReq.Username),
		PasswordHash: string(hashedPassword),
	}

	if err := env.store.CreateUser(&model); err != nil {
		http.Error(responseWriter,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
}
