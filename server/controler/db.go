package controler

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/reactivex/rxgo/v2"

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

func (dbStore *DbStore) GetUserByUsername(username string) rxgo.Observable {
	return rxgo.Create([]rxgo.Producer{func(_ context.Context, next chan<- rxgo.Item) {
		var result model.User

		err := dbStore.db.Where("username = ?", username).First(&result).Error
		if err != nil {
			next <- rxgo.Error(err)
		}
		next <- rxgo.Of(&result)
	}})

}

func (dbStore *DbStore) IsUserPresent(username string) (bool, error) {
	return false, nil
	// user, err := dbStore.GetUserByUsername(username)
	// if err != nil {
		// return false, err
	// }

	// return user != nil, nil
}

func (dbStore *DbStore) CreateUser(newUser *model.User) error {
	return dbStore.db.Create(newUser).Error
}
