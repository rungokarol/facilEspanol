package controler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TODO: move to app config
var jwtKey = []byte("facil_espanol_jwt_key")
var expiryHours = time.Duration(24)

type claims struct {
	username string
	jwt.StandardClaims
}

func createJwt(username string) (string, error) {
	expirationTime := time.Now().Add(expiryHours * time.Hour)
	claims := &claims{
		username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
