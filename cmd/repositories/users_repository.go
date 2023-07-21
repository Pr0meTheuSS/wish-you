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

package repository

import (
	"errors"

	"github.com/lib/pq"
)

type UserRepository interface {
	InsertUser(user RepositoryUser) error
	GetUserByEmail(email string) (*RepositoryUser, error)
	AuthenticateUser(email string, password string) (bool, error)
}

type RepositoryUser struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Confirmed bool   `db:"confirmed"`
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

		return errors.New(err.Error())
	}

	return nil
}
