package main

/*
 * Project: I-wish-you-app
 * Created Date: Wednesday, July 5th 2023, 7:00:28 pm
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
	"fmt"
	"log"
	repository "main/cmd/repositories"
	"main/cmd/server"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO: Выделить абстракции и перенести их на подходящий уровень проекта
// TODO: Добавить изящное завершение приложения по сигналам прерывания
// TODO: Добавить логирование проекта

type Configurations struct {
	addr        string
	port        int
	serviceName string
}

type App struct {
	configurations Configurations
}

type Runnable interface {
	run() error
}

type Message struct {
	Text string `json:"message"`
}

func (a App) run() error {
	// Открытие соединения с базой данных для проверки корректности работы связи с бд
	db, err := sqlx.Connect("postgres", repository.GetDBDSN())
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer db.Close()

	server := server.NewServer(a.configurations.addr, a.configurations.port)
	log.Printf(
		"Service %s is running, bro, server is listening port %d\n",
		a.configurations.serviceName,
		a.configurations.port)

	return server.Run()
}

func loadConfigurations(configurations *Configurations) error {
	// TODO: Реализовать чтение конфигураций из внешнего файла или из окружения
	*configurations = Configurations{
		addr:        "127.0.0.1",
		port:        6969,
		serviceName: "wish-you",
	}

	return nil
}

func main() {
	appConfigurations := Configurations{}
	if loadConfigurations(&appConfigurations) != nil {
		// TODO: Обработать ошибки
	}

	wishYouApp := App{appConfigurations}
	if err := wishYouApp.run(); err != nil {
		// TODO: Обработать ошибки
		fmt.Print(err)
	}
}
