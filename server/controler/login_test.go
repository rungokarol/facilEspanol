package controler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rungokarol/facilEspanol/mocks"
	"github.com/rungokarol/facilEspanol/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

var username string = "foo"

type LoginReqTestSuite struct {
	suite.Suite
	env       *Env
	storeMock *mocks.IDataStore
	rr        *httptest.ResponseRecorder
	handler   http.Handler
}

func (suite *LoginReqTestSuite) SetupTest() {
	suite.storeMock = &mocks.IDataStore{}
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

func (suite *LoginReqTestSuite) TestRejectIfUserNotFoundInDataStore() {
	jsonBody, err := json.Marshal(loginReq{Username: username, Password: "bar"})
	assert.Nil(suite.T(), err)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonBody))
	assert.Nil(suite.T(), err)

	suite.storeMock.On("GetUserByUsername", username).Return(nil, nil)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusNotFound, suite.rr.Code)
}

func (suite *LoginReqTestSuite) TestRejectIfErrorOccursDuringExtractingUserFromStore() {
	jsonBody, err := json.Marshal(loginReq{Username: username, Password: "bar"})
	assert.Nil(suite.T(), err)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonBody))
	assert.Nil(suite.T(), err)

	suite.storeMock.On("GetUserByUsername", username).
		Return(nil, errors.New("DEADBEEF"))

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.rr.Code)
}

func (suite *LoginReqTestSuite) TestRejectIfPasswordIsIncorrect() {
	jsonBody, err := json.Marshal(loginReq{Username: username, Password: "bar"})
	assert.Nil(suite.T(), err)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonBody))
	assert.Nil(suite.T(), err)

	suite.storeMock.On("GetUserByUsername", username).
		Return(&model.User{Username: username, PasswordHash: "no_match"}, nil)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusForbidden, suite.rr.Code)
}

func (suite *LoginReqTestSuite) TestAcceptIfUserExistsInStoreAndPassowrdMatches() {
	correctHash, err := bcrypt.GenerateFromPassword([]byte("correct"), bcrypt.MinCost)
	assert.Nil(suite.T(), err)
	jsonBody, err := json.Marshal(loginReq{Username: username, Password: "correct"})
	assert.Nil(suite.T(), err)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonBody))
	assert.Nil(suite.T(), err)

	suite.storeMock.On("GetUserByUsername", username).
		Return(&model.User{Username: username, PasswordHash: string(correctHash)}, nil)

	suite.handler.ServeHTTP(suite.rr, req)
	assert.Equal(suite.T(), http.StatusOK, suite.rr.Code)
}
