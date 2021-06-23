package controler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rungokarol/facilEspanol/model"
	"golang.org/x/crypto/bcrypt"
)

type registerReq struct {
	Username string
	Password string
	Email    string
}

func (env *Env) Register(responseWriter http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var registerReq registerReq
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(registerReq.Username) < minLength || len(registerReq.Password) < minLength {
		http.Error(responseWriter,
			"Username or password too short",
			http.StatusBadRequest)
		return
	}

	emailInUse, err := env.store.EmailAlreadyInUse(registerReq.Email)
	if emailInUse {
		http.Error(responseWriter,
			"Email already in use",
			http.StatusBadRequest)
		return
	}

	username := strings.ToLower(registerReq.Username)
	isPresent, err := env.store.IsUserPresent(username)
	if err != nil {
		//check if good status maybe unauthorized?
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	} else if isPresent {
		http.Error(responseWriter,
			"User with given username already exists",
			http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	newUser := model.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}

	if err := env.store.CreateUser(&newUser); err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
}
