package main

import (
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/rungokarol/facilEspanol/controler"
)

func main() {
	env := controler.CreateEnv()
	defer env.Close() //make it private

	http.HandleFunc("/", env.DefaultRoot)
	http.HandleFunc("/user/login", env.Login)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
