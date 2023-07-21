package repository

import (
	"log"
	"main/cmd/server/hash"

	"os"

	"github.com/jmoiron/sqlx"
)

/*
 * Project: YourProjectName
 * Created Date: Wednesday, July 19th 2023, 10:08:13 pm
 * Author: Olimpiev Y. Y.
 * -----
 * Last Modified:  yr.olimpiev@gmail.com
 * Modified By: Olimpiev Y. Y.
 * -----
 * Copyright (c) 2023 NSU
 *
 * -----
 */

type SqlxUsersRepository struct {
	DB *sqlx.DB
}

func GetDBDSN() string {
	return os.Getenv("DB_DSN")
}

func (repo *SqlxUsersRepository) InsertUser(repositoryUser RepositoryUser) error {
	log.Printf("Insert user %v", repositoryUser)
	// Открытие соединения с базой данных
	var err error
	repo.DB, err = sqlx.Connect("postgres", GetDBDSN())
	if err != nil {
		log.Println(err.Error())
	}

	defer repo.DB.Close()

	log.Printf("Exec query for insert user %v", repositoryUser)
	_, err = repo.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		repositoryUser.Name,
		repositoryUser.Email,
		repositoryUser.Password)

	return handleDBErrors(err)
}

func (repo *SqlxUsersRepository) GetUserByEmail(email string) (*RepositoryUser, error) {
	var err error
	repo.DB, err = sqlx.Connect("postgres", GetDBDSN())
	if err != nil {
		log.Println(err.Error())
	}

	defer repo.DB.Close()
	users := []RepositoryUser{}

	err = repo.DB.Select(&users, "SELECT * FROM users WHERE (email=$1)", email)
	if len(users) != 1 {
		return nil, err
	}
	// TODO: добавить проверку ошибок noRows и MultiRows
	return &users[0], handleDBErrors(err)
}

func (repo *SqlxUsersRepository) AuthenticateUser(email string, password string) (bool, error) {
	users := []RepositoryUser{}
	var err error
	repo.DB, err = sqlx.Connect("postgres", GetDBDSN())
	if err != nil {
		log.Println(err.Error())
	}

	defer repo.DB.Close()
	log.Printf("user email : %s\n", email)
	err = repo.DB.Select(&users, "SELECT * FROM users WHERE users.email = $1;", email)
	if err != nil {
		log.Println(err.Error())
	}

	if len(users) != 1 {
		log.Println(len(users))
		return false, err
	}

	// TODO: добавить проверку ошибок noRows и MultiRows
	log.Println(hash.VerfyPassword(password, users[0].Password))
	log.Println(handleDBErrors(err))
	return len(users) == 1 && hash.VerfyPassword(password, users[0].Password), handleDBErrors(err)
}
