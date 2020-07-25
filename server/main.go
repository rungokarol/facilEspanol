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

func initDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=facilEspanolUser dbname=facilEspanolDb password=facilEspanolPass sslmode=disable")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}

	log.Println("Connected to database!")

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	return db
}

func main() {
	db := initDb()
	defer db.Close()

	http.HandleFunc("/", handler)
	http.HandleFunc("/login", login.LoginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
