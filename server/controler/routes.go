package controler

import (
	"fmt"
	"log"
	"net/http"
)

func (env *Env) DefaultRoot(responseWriter http.ResponseWriter, r *http.Request) {
	log.Println("request received")
	fmt.Fprintf(responseWriter, "Hello %s!", r.URL.Path[1:])
}

func (env *Env) Login(responseWriter http.ResponseWriter, r *http.Request) {
	log.Println("czesc i czolem!")
	fmt.Fprintf(responseWriter, "user login")
}
