package helper

import (
	"github.com/goupp-backend/model"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
    "time"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWTToken(user model.User) (string, error){
	tokenTTL, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_TTL"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"id": user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}
