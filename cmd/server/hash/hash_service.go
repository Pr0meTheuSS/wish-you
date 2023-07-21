package hash

/*
 * Project: YourProjectName
 * Created Date: Friday, July 21st 2023, 12:47:25 am
 * Author: Olimpiev Y. Y.
 * -----
 * Last Modified:  yr.olimpiev@gmail.com
 * Modified By: Olimpiev Y. Y.
 * -----
 * Copyright (c) 2023 NSU
 *
 * -----
 */

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GetHash(value string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func VerfyPassword(password string, expected_hash string) bool {
	log.Printf("password and hash %s, %s", password, expected_hash)

	isPasswordVerified := bcrypt.CompareHashAndPassword([]byte(expected_hash), []byte(password))

	log.Printf("error in verify function %v", isPasswordVerified)
	return isPasswordVerified == nil
}
