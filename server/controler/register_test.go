package controler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rungokarol/facilEspanol/mocks"
	"github.com/rungokarol/facilEspanol/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RegisterTestSuite struct {
	suite.Suite
	env       *Env
	storeMock *mocks.IDataStore
	rr        *httptest.ResponseRecorder
	handler   http.Handler
}

func (suite *RegisterTestSuite) SetupTest() {
	suite.storeMock = &mocks.IDataStore{}
	suite.env = CreateEnv(suite.storeMock)
	suite.rr = httptest.NewRecorder()
	suite.handler = http.HandlerFunc(suite.env.Register)
}

func TestRegister(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}

func (suite *RegisterTestSuite) TestRejectWithNotPostMethod() {
	req, err := http.NewRequest("GET", "/user/register", nil)
	assert.Nil(suite.T(), err)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusMethodNotAllowed, suite.rr.Code)
}

func (suite *RegisterTestSuite) TestRejectWhenBodyIsNotJson() {
	notJsonBody := ioutil.NopCloser(bytes.NewBufferString("Hello World"))
	req, err := http.NewRequest("POST", "/user/register", notJsonBody)
	assert.Nil(suite.T(), err)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.rr.Code)
}

func (suite *RegisterTestSuite) TestRejectIfUsernameTooShort() {
	jsonBody, err := json.Marshal(registerReq{Username: "ja", Password: "bar"})
	assert.Nil(suite.T(), err)
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonBody))
	assert.Nil(suite.T(), err)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.rr.Code)
}

func (suite *RegisterTestSuite) TestRejectIfPasswordTooShort() {
	jsonBody, err := json.Marshal(registerReq{Username: "foo", Password: "ja"})
	assert.Nil(suite.T(), err)
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonBody))
	assert.Nil(suite.T(), err)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.rr.Code)
}

func (suite *RegisterTestSuite) TestRejectIfUserAlreadyExists() {
	username := "foo"
	jsonBody, err := json.Marshal(registerReq{Username: username, Password: "bar"})
	assert.Nil(suite.T(), err)
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonBody))
	assert.Nil(suite.T(), err)

	suite.storeMock.
		On("IsUserPresent", username).Return(true, nil).
		On("EmailAlreadyInUse", mock.Anything).Return(false, nil)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.rr.Code)
}

func (suite *RegisterTestSuite) TestAcceptRegistration() {
	username := "foo"
	jsonBody, err := json.Marshal(registerReq{Username: username, Password: "bar"})
	assert.Nil(suite.T(), err)
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonBody))
	assert.Nil(suite.T(), err)

	suite.storeMock.
		On("IsUserPresent", username).Return(false, nil).
		On("EmailAlreadyInUse", mock.Anything).Return(false, nil).
		On("CreateUser", mock.Anything).Return(nil)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusOK, suite.rr.Code)
	suite.storeMock.AssertCalled(suite.T(), "CreateUser", mock.MatchedBy(func(user *model.User) bool { return user.Username == username }))

}

func (suite *RegisterTestSuite) TestRejectIfEmailAlreadyExists() {
	username := "foo"
	jsonBody, err := json.Marshal(registerReq{Username: username, Password: "bar", Email: "dummy@email.com"})
	assert.Nil(suite.T(), err)
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer(jsonBody))
	assert.Nil(suite.T(), err)

	suite.storeMock.
		On("EmailAlreadyInUse", mock.Anything).Return(true, nil)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.rr.Code)
}
