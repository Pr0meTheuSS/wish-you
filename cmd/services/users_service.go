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

type UsersService interface {
	RegisterUser(user ServiceUser) error
	AuthenticateUser(authRequest AuthenticateUserRequest) (bool, error)
}

type ServiceUser struct {
	Name     string
	Email    string
	Password string
}

type AuthenticateUserRequest struct {
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
