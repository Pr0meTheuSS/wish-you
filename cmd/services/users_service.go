package service

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

import (
	"log"
	repository "main/cmd/repositories"
)

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
	if err := repository.InsertUser(repository.RepositoryUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}); err != nil {
		return err
	}

	if err := SendConfirmMail(user.Email); err != nil {
		log.Printf("Ошибка отправки сообщения с подтверждением")
		return err
	}

	return nil
}
