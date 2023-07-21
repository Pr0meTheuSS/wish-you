package internal

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/*
 * Project: YourProjectName
 * Created Date: Friday, July 21st 2023, 12:14:31 am
 * Author: Olimpiev Y. Y.
 * -----
 * Last Modified:  yr.olimpiev@gmail.com
 * Modified By: Olimpiev Y. Y.
 * -----
 * Copyright (c) 2023 NSU
 *
 * -----
 */

var secretKey = []byte("ваш_секретный_ключ")

func GenerateToken(userEmail string, userPassword string) (string, error) {
	// Создание нового токена
	token := jwt.New(jwt.SigningMethodHS256)

	// Создание клейма (payload) с информацией о пользователе
	claims := token.Claims.(jwt.MapClaims)
	claims["useremail"] = userEmail
	claims["userpassword"] = userPassword
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Время жизни токена (здесь 24 часа)

	// Подпись токена с использованием секретного ключа
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что используется правильный алгоритм подписи (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetJwtClaims(token *jwt.Token) (JWTClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		log.Printf("Token %v is valid\n", token)
		// Доступ к payload (заявкам)
		log.Println("Email:", claims["useremail"])
		log.Println("Password:", claims["userpassword"])
		log.Println("Expiration Time:", claims["exp"])
	} else {
		log.Println("Token is invalid.")
	}

	return JWTClaims{
		Email:      claims["useremail"].(string),
		Password:   claims["userpassword"].(string),
		Expiration: claims["exp"].(float64),
	}, nil
}

type JWTClaims struct {
	Email      string
	Password   string
	Expiration float64
}

func IsJWTExpired(claims JWTClaims) bool {
	return int64(claims.Expiration) < time.Now().Unix()
}
