package tokens

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

//https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac

var secretKey []byte = []byte("jobless")

// Access токен тип JWT, алгоритм SHA512,
func NewJWTToken(mail, ip string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ip":   ip,
		"mail": mail,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (string, string, error) {
	keyfunc := func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method %v", t.Header["alg"])
		}

		return secretKey, nil
	}

	token, err := jwt.Parse(tokenString, keyfunc)
	if err != nil {
		log.Println("faild to parse", err)
		return "", "", err
	}

	ip := ""
	mail := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if ipVal, ok := claims["ip"].(string); ok {
			ip = ipVal
		} else {
			return "", "", fmt.Errorf("could not get ip")
		}
		if mailVal, ok := claims["mail"].(string); ok {
			mail = mailVal
		} else {
			return "", "", fmt.Errorf("could not get mail")
		}
	} else {
		return "", "", fmt.Errorf("the token is not valid")
	}

	return mail, ip, nil
}
