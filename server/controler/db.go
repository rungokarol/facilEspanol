package controler

import (
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

	err := dbStore.db.Where("username = ?", username).First(&result).Error

	if err != nil && gorm.IsRecordNotFoundError(err) {
		return result, err
	} else if err != nil {
		log.Println("Error fetching from database: ", err)
		return result, err
	}

	return result, nil
}
