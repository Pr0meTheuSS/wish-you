package serviceuser

import (
	repository "main/cmd/repositories"
)

/*
 * Project: I-wish-you
 * Created Date: Thursday, July 13th 2023, 11:14:21 pm
 * Author: Olimpiev Y. Y.
 * -----
 * Last Modified:  yr.olimpiev@gmail.com
 * Modified By: Olimpiev Y. Y.
 * -----
 * Copyright (c) 2023 NSU
 *
 * -----
 */

type ServiceUser struct {
	Name     string
	Email    string
	Password string
}

func NewUser(name string, email string, password string) ServiceUser {
	return ServiceUser{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func RegisterUser(user ServiceUser) error {
	// TODO: обработать ошибки и смаппить ошибку репозитория в ошибку сервиса
	return repository.InsertUser(repository.RepositoryUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
}
