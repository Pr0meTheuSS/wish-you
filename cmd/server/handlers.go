package server

/*
 * Project: I-wish-you
 * Created Date: Sunday, July 9th 2023, 10:58:30 am
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

	"github.com/gin-gonic/gin"
)

func loginGetHandler(ctx *gin.Context) {
	body, err := ioutil.ReadFile("cmd/pages/login-page/login.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

func rootPageGetHandler(ctx *gin.Context) {
	body, err := ioutil.ReadFile("cmd/pages/main-page/main.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

type Message struct {
	Text string `json:"message"`
}

func sendMessagePostHandler(ctx *gin.Context) {
	// Получение тела запроса в виде структуры Message
	var message Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// TODO: Убрать на релизе этот вывод
	// Использование текста сообщения
	fmt.Println("Текст сообщения:", message.Text)

	// Ваша логика обработки сообщения

	// Ответ сервера
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func loginPostHandler(ctx *gin.Context) {
	body, err := ioutil.ReadFile("cmd/pages/login-page/success.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// TODO: Убрать на релизе этот вывод
	fmt.Printf("Handle post request for /login\n [username]: %s\t [password]: %s\n", username, password)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}
