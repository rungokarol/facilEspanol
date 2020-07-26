package controler

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/rungokarol/facilEspanol/model"
)

type Env struct {
	db *gorm.DB
}

func CreateEnv() *Env {
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

func (env *Env) Close() {
	env.db.Close()
}
