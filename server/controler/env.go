package controler

import "github.com/rungokarol/facilEspanol/model"

type IDataStore interface {
	GetUserByUsername(string) (*model.User, error)
	IsUserPresent(string) (bool, error)
	CreateUser(*model.User) error
	EmailAlreadyInUse(string) (bool, error)
}

type Env struct {
	store IDataStore
	//other pointers will go here
}

func CreateEnv(store IDataStore) *Env {
	return &Env{store: store}
}
