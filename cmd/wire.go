//go:build wireinject
// +build wireinject

// wire.go
package wire

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	repository "main/cmd/repositories"
	server "main/cmd/server"
	service "main/cmd/services"
)

func NewRealUserRepository() *repository.SqlxUsersRepository {
	wire.Build(ProvideRealUserRepository)
	return nil
}

func ProvideRealUserRepository() *repository.SqlxUsersRepository {
	return &repository.SqlxUsersRepository{
		DB: &sqlx.DB{},
	}
}

func NewRealUsersService() *service.UsersService {
	wire.Build(ProvideRealUsersService)
	return nil
}

func ProvideRealUsersService() *service.UsersService {
	return NewRealUsersService()
}

func NewRealHandler() *server.Handler {
	wire.Build(ProvideRealHander)
	return nil
}

func ProvideRealHander() *server.Handler {
	return &server.Handler{
		UsersService: *NewRealUsersService(),
	}
}

// func NewMockUserRepository() *MockUserRepository {
// 	wire.Build(ProvideMockUserRepository)
// 	return nil
// }

// func ProvideMockUserRepository() *MockUserRepository {
// 	return &MockUserRepository{}
// }
