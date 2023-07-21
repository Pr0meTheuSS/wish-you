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

type RealUsersService struct {
	UsersRepo repository.UserRepository
}

func (s *RealUsersService) RegisterUser(user ServiceUser) error {

	if err := s.UsersRepo.InsertUser(repository.RepositoryUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}); err != nil {
		return err
	}

	// TODO: прокинуть тело письма подтверждения почты
	if err := SendConfirmMail(user.Email, "Привет, спасибо за регистрацию в IWY!"); err != nil {
		log.Printf("Ошибка отправки сообщения с подтверждением")
		return err
	}

	return nil
}
func (s *RealUsersService) AuthenticateUser(authRequest AuthenticateUserRequest) (bool, error) {
	return s.UsersRepo.AuthenticateUser(authRequest.Email, authRequest.Password)

}
