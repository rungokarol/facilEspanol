package controler

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/reactivex/rxgo/v2"
	"github.com/rungokarol/facilEspanol/model"
	"golang.org/x/crypto/bcrypt"
)

type loginReq struct {
	Username string
	Password string
}

type loginResp struct {
	Token string `json:"token"`
}

var minLength = 3

func HttpBodyBytesOb(r *http.Request) rxgo.Observable {
	return rxgo.Create([]rxgo.Producer{func(_ context.Context, next chan<- rxgo.Item) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			next <- rxgo.Error(err)
			return
		}
		next <- rxgo.Of(bytes)
	}})
}

func HttpRequestOb(r *http.Request) rxgo.Observable {
	ch := make(chan rxgo.Item)
	ch <- rxgo.Of(r.Method)
	return rxgo.FromChannel(ch)
}

func ErrorIfHttpMethodIsNot(method string) rxgo.Func {
	return func(_ context.Context, item interface{}) (interface{}, error) {
		request := item.(*http.Request)
		if request.Method != method {
			// TODO: pass http code
			return nil, errors.New("asdas")
		}
		return request, nil
	}
}

func ErrorIfUserNotFound() rxgo.Func {
	return func(_ context.Context, item interface{}) (interface{}, error) {
		user := item.(*model.User)
		if user == nil {
			return nil, errors.New("User not found")
		}
		return user, nil
	}
}

func ErrorIfPasswordHashDoesNotMatch(requestPassword string) rxgo.Func {
	return func(_ context.Context, item interface{}) (interface{}, error) {
		user := item.(*model.User)
		if err := bcrypt.CompareHashAndPassword(
			[]byte(user.PasswordHash), []byte(requestPassword)) ; err != nil {
			return nil, err
		}
		return user, nil
	}
}

func CreateJwtToken() rxgo.Func {
	return func(_ context.Context, item interface{}) (interface{}, error) {
		user := item.(*model.User)
		return createJwt(user.Username)
	}
}

func (env *Env) Login(responseWriter http.ResponseWriter, r *http.Request) {
	obs := HttpRequestOb(r).Map(
		ErrorIfHttpMethodIsNot(http.MethodPost),
	).FlatMap(func(item rxgo.Item) rxgo.Observable {
		request := item.V.(*http.Request)
		return HttpBodyBytesOb(request)
	}).Unmarshal(
		json.Unmarshal,
		func() interface{} { return &loginReq{} },
	).FlatMap(func(item rxgo.Item) rxgo.Observable {
		loginReq := item.V.(loginReq)
		return env.store.GetUserByUsername(strings.ToLower(loginReq.Username),
			).Map(ErrorIfUserNotFound(),
			).Map(ErrorIfPasswordHashDoesNotMatch(loginReq.Password))
	}).Map(CreateJwtToken(),
	).Marshal(json.Marshal)

	obs.DoOnNext(func(item interface{}) {
		jsonRes := item.([]byte)
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Write(jsonRes)
	})
	obs.DoOnError(func(err error) {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
	})

}
