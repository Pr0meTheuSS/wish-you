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
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// TODO: replace const string to getenv
func GetDBDSN() string {
	return "postgres://postgres:password@localhost:5431/postgres?sslmode=disable"
}

var DB *sqlx.DB

type RepositoryUser struct {
	Name     string
	Email    string
	Password string
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
	// TODO: размаршалить ошибку, то есть
	// Понять что произошло на уровне запроса и вернуть ответ на уровне репози
	return handleDBErrors(err)
}

func handleDBErrors(err error) error {
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Constraint == "users_email_key" && pqErr.Code.Name() == "unique_violation" {
				// Обработка ошибки дублирования уникального ключа
				fmt.Println("Пользователь с таким email уже существует")
				return errors.New("constraint error: user already exists")
			}
		}
		return errors.New("unknown error")
	}

	return nil
}
