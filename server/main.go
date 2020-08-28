package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

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

	mux := http.NewServeMux()
	mux.HandleFunc("/user/login", env.Login)
	mux.HandleFunc("/user/register", env.Register)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	log.Fatal(http.ListenAndServe(":8080", cors.Handler(mux)))
}
