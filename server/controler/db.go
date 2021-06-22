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

func (dbStore *DbStore) GetUserByUsername(username string) (*model.User, error) {
	var result model.User

	err := dbStore.db.Where("username = ?", username).First(&result).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		log.Println("Error fetching from database: ", err)
		return nil, err
	} else if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}

	return &result, nil
}

func (dbStore *DbStore) IsUserPresent(username string) (bool, error) {
	user, err := dbStore.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	return user != nil, nil
}

func (dbStore *DbStore) CreateUser(newUser *model.User) error {
	return dbStore.db.Create(newUser).Error
}

func (dbStore *DbStore) EmailAlreadyInUse(email string) (bool, error) {
	var users []model.User
	result := dbStore.db.Where("email <> ?", email).Find(&users)
	if result.RowsAffected > 0 {
		return true, result.Error
	} else {
		return false, result.Error
	}
}
