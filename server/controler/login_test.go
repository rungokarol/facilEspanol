package controler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rungokarol/facilEspanol/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

type LoginReqTestSuite struct {
	suite.Suite
	env       *Env
	storeMock *storeMock
	rr        *httptest.ResponseRecorder
	handler   http.Handler
}

func (suite *LoginReqTestSuite) SetupTest() {
	suite.storeMock = &storeMock{}
	suite.env = CreateEnv(suite.storeMock)
	suite.rr = httptest.NewRecorder()
	suite.handler = http.HandlerFunc(suite.env.Login)
}

func TestLoginReq(t *testing.T) {
	suite.Run(t, new(LoginReqTestSuite))
}

func (suite *LoginReqTestSuite) TestRejectWithNotPostMethod() {
	req, err := http.NewRequest("GET", "/user/login", nil)
	assert.Nil(suite.T(), err)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusMethodNotAllowed, suite.rr.Code)
}

func (suite *LoginReqTestSuite) TestRejectWhenBodyIsNotJson() {
	notJsonBody := ioutil.NopCloser(bytes.NewBufferString("Hello World"))
	req, err := http.NewRequest("POST", "/user/login", notJsonBody)
	assert.Nil(suite.T(), err)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.rr.Code)
}

// TODO
// 1. create "testify" mock
