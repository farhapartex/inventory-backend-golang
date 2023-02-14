package helper

import (
	"github.com/goupp-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"fmt"
	"strings"
	"errors"
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


func GetTokenFromRequet(context *gin.Context) string{
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
        return splitToken[1]
    }
    return ""
}

func GetToken(context *gin.Context) (*jwt.Token, error){
	tokenString := GetTokenFromRequet(context)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return privateKey, nil
    })
    return token, err
}

func ValidateJWT(context *gin.Context) error {
    token, err := GetToken(context)
    if err != nil {
        return err
    }
    _, ok := token.Claims.(jwt.MapClaims)
    if ok && token.Valid {
        return nil
    }
    return errors.New("invalid token provided")
}


func CurrentUser(context *gin.Context) (model.User, error) {
    err := ValidateJWT(context)
    if err != nil {
        return model.User{}, err
    }
    token, _ := GetToken(context)
    claims, _ := token.Claims.(jwt.MapClaims)
    userId := uint(claims["id"].(float64))

    user, err := model.FindUserById(userId)
    if err != nil {
        return model.User{}, err
    }
    return user, nil
}

