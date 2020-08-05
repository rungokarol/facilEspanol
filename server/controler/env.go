package controler

import (
	"github.com/reactivex/rxgo/v2"
	"github.com/rungokarol/facilEspanol/model"
)

type IDataStore interface {
	GetUserByUsername(string) rxgo.Observable
	IsUserPresent(string) (bool, error)
	CreateUser(*model.User) error
}

type Env struct {
	store IDataStore
	//other pointers will go here
}

func CreateEnv(store IDataStore) *Env {
	return &Env{store: store}
}
