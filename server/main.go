package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/rungokarol/facilEspanol/login"
	model "github.com/rungokarol/facilEspanol/model/db"
)

func handler(responseWriter http.ResponseWriter, r *http.Request) {
	log.Println("request received")

	fmt.Fprintf(responseWriter, "Hello %s!", r.URL.Path[1:])
}

func createEnv() *Env {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=facilEspanolUser dbname=facilEspanolDb password=facilEspanolPass sslmode=disable")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}

	log.Println("Connected to database!")

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	return &Env{db}
}

type Env struct {
	db *gorm.DB
}

func (env *Env) loginHandler(responseWriter http.ResponseWriter, r *http.Request) {
	login.HandleRequest(responseWriter)
}

func main() {
	env := createEnv()
	defer env.db.Close() //make it private

	http.HandleFunc("/", handler)
	http.HandleFunc("/user/login", env.loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
