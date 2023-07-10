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
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: Выделить абстракции и перенести их на подходящий уровень проекта
// TODO: Добавить изящное завершение приложения по сигналам прерывания
// TODO: Добавить логирование проекта

type Configurations struct {
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
	fmt.Printf(
		"Service %s is running, bro, server is listening port %d\n",
		a.configurations.serviceName,
		a.configurations.port)

	router := gin.Default()

	router.GET("/login", func(ctx *gin.Context) {
		body, err := ioutil.ReadFile("cmd/pages/login-page/login.html")
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
	})

	router.GET("/", func(ctx *gin.Context) {
		body, err := ioutil.ReadFile("cmd/pages/main-page/main.html")
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
	})

	router.POST("/send-message", func(ctx *gin.Context) {
		// Получение тела запроса в виде структуры Message
		var message Message
		if err := ctx.ShouldBindJSON(&message); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		// Использование текста сообщения
		fmt.Println("Текст сообщения:", message.Text)

		// Ваша логика обработки сообщения

		// Ответ сервера
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	router.POST("/login", func(ctx *gin.Context) {
		body, err := ioutil.ReadFile("cmd/pages/login-page/success.html")
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		// TODO: Убрать на релизе этот вывод
		fmt.Printf("Handle post request for /login\n [username]: %s\t [password]: %s\n", username, password)

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
	})

	return router.Run(":" + strconv.Itoa(a.configurations.port))
}

func loadConfigurations(configurations *Configurations) error {
	// TODO: Реализовать чтение конфигураций из внешнего файла
	*configurations = Configurations{
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
