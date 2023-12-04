package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("dei-secret-key")

func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   string(userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// package helpers

// import (
// 	"errors"
// 	"github.com/dgrijalva/jwt-go"
// 	"time"
// )

// var jwtKey = []byte("your_secret_key") // Replace with your secret key

// type Claims struct {
// 	UserID uint `json:"userId"`
// 	jwt.StandardClaims
// }

// func GenerateToken(userID uint) (string, error) {
// 	claims := &Claims{
// 		UserID: userID,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
// 			IssuedAt:  time.Now().Unix(),
// 			Issuer:    "your_issuer", // Replace with your issuer
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedToken, nil
// }

// func VerifyToken(tokenString string) (*jwt.Token, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, errors.New("invalid token")
// 	}

// 	return token, nil
// }
