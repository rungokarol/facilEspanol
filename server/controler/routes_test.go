package controler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rungokarol/facilEspanol/model"
)

type storeMock struct {
}

func (sm *storeMock) GetUserByUsername(string) (*model.User, error) {
	return nil, nil
}
func (sm *storeMock) IsUserPresent(string) (bool, error) {
	return false, nil
}
func (sm *storeMock) CreateUser(*model.User) error {
	return nil
}

func TestShouldRejectUserLoginReqWithNotPostMethod(t *testing.T) {
	storeMock := &storeMock{}
	env := CreateEnv(storeMock)

	req, err := http.NewRequest("GET", "/user/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.Login)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 405 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
