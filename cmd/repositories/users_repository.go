package repository

// TODO: протестировать работу слоя репозитория

/*
 * Project: I-wish-you
 * Created Date: Thursday, July 13th 2023, 10:43:01 pm
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
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// TODO: replace const string to getenv
func GetDBDSN() string {
	return "postgres://postgres:password@localhost:5431/postgres?sslmode=disable"
}

// TODO: описать интерфейс базы данных
var DB *sqlx.DB

type RepositoryUser struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func InsertUser(repositoryUser RepositoryUser) error {
	log.Printf("Insert user %v", repositoryUser)
	// Открытие соединения с базой данных
	db, err := sqlx.Connect("postgres", GetDBDSN())
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer db.Close()

	log.Printf("Exec query for insert user %v", repositoryUser)
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		repositoryUser.Name,
		repositoryUser.Email,
		repositoryUser.Password)

	return handleDBErrors(err)
}

func handleDBErrors(err error) error {
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Constraint == "users_email_key" && pqErr.Code.Name() == "unique_violation" {
				// Обработка ошибки дублирования уникального ключа
				return errors.New("constraint error: user already exists")
			}
		}

		return errors.New("unknown error")
	}

	return nil
}

func GetUser(repositoryUser RepositoryUser) (RepositoryUser, error) {
	log.Printf("Insert user %v", repositoryUser)
	// Открытие соединения с базой данных
	db, err := sqlx.Connect("postgres", GetDBDSN())
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer db.Close()

	log.Printf("Exec query for insert user %v", repositoryUser)
	users := []RepositoryUser{}

	err = db.Select(&users, "SELECT id, username, email FROM users WHERE (email=$1 AND password=$2)", repositoryUser.Email, repositoryUser.Password)
	// TODO: добавить проверку ошибок noRows и MultiRows
	return users[0], handleDBErrors(err)
}
