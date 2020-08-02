package main

import (
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/rungokarol/facilEspanol/controler"
)

func main() {
	store, err := controler.OpenDB()
	if err != nil {
		panic("Could not connect to database!")
	}
	defer store.Close()

	env := controler.CreateEnv(store)

	http.HandleFunc("/", env.DefaultRoot)
	http.HandleFunc("/user/login", env.Login)
	http.HandleFunc("/user/register", env.Register)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
