package common

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func GenToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	return token.SignedString([]byte("2108a"))
}

func CheckToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		return []byte("2108a"), nil
	})
	if err != nil {
		return 0, err
	}

	tokenMap := token.Claims.(jwt.MapClaims)
	fmt.Println("***********tokenMap")
	fmt.Println(tokenMap)
	if _, ok := tokenMap["user_id"]; !ok {
		return 0, fmt.Errorf("token invalid")
	}
	return int64(tokenMap["user_id"].(float64)), err
}
