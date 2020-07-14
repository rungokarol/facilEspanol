package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "github.com/rungokarol/facilEspanol/login"

)

type Product struct {
  gorm.Model
  Code string
  Name string
  Price uint
}

func handler(responseWriter http.ResponseWriter, r *http.Request) {
    log.Println("request received")

    fmt.Fprintf(responseWriter, "Hello %s!", r.URL.Path[1:])
}

func main() {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=facilespanoluser dbname=facilespanoldb password=facilEspanolPass  sslmode=disable")
    if err != nil {
      log.Println(err)
      panic("failed to connect database")
    } else {
      log.Println("Connected to database!")
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&Product{})

    // Create
    db.Create(&Product{Code: "L1212", Price: 1000})


    http.HandleFunc("/", handler)
    http.HandleFunc("/login", login.LoginHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

