package controler

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/rungokarol/facilEspanol/model"
)

type DbStore struct {
	db *gorm.DB
}

func OpenDB() (*DbStore, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=facilEspanolUser dbname=facilEspanolDb password=facilEspanolPass sslmode=disable")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Connected to database!")

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	return &DbStore{db: db}, nil
}

func (db *DbStore) Close() {
	db.Close()
}

func (dbStore *DbStore) GetUserByUsername(username string) (model.User, error) {
	var result model.User
	if dbStore.db.Where("username = ?", username).First(&result).RecordNotFound() {
		return result, errors.New("User not found!") //try another way to get error
	}

	return result, nil
}
