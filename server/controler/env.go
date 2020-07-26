package controler

type IDataStore interface {
}

type Env struct {
	store IDataStore
	//other pointers will go here
}

func CreateEnv(store IDataStore) *Env {
	return &Env{store: store}
}
