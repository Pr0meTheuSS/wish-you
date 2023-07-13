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
	"log"
	serviceuser "main/cmd/services/ServiceUser"

	"github.com/jmoiron/sqlx"
)

// TODO: replace const string to getenv
func getDBDSN() string {
	return "postgres://postgres:password@localhost:5431/postgres?sslmode=disable"
}

var DB *sqlx.DB

func init() {
	DB, err := sqlx.Open("postgres", getDBDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
}

type RepositoryUser struct {
	Name     string
	Email    string
	Password string
}

func IsnertUser(serviceUser serviceuser.ServiceUser) error {
	repositoryUser := RepositoryUser{
		Name:     serviceUser.Name,
		Email:    serviceUser.Email,
		Password: serviceUser.Password,
	}
	// TODO: Добавить миграцию базы для увеличения метаинформации о пользователе
	//DB.Exec("INSERT INTO users (name, email, password, createdat) VALUES (?, ?, ?)")
	_, err := DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		repositoryUser.Name,
		repositoryUser.Email,
		repositoryUser.Password)
	// TODO: размаршалить ошибку, то есть
	// Понять что произошло на уровне запроса и вернуть ответ на уровне репозитория
	return err
}
