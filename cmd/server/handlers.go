package server

// REFACTOR:  перенести полезную нагрузку работы хэндлеров на уровень сервисов
// REFACTOR:  перенести используемые типы в отдельный файл

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
	"encoding/json"
	"fmt"
	"log"
	serviceuser "main/cmd/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var handlers = map[string]gin.HandlerFunc{
	"rootPageGetHandler":      rootPageGetHandler,
	"loginGetHandler":         loginGetHandler,
	"loginPostHandler":        loginPostHandler,
	"sendMessagePostHandler":  sendMessagePostHandler,
	"signinPostHandler":       signinPostHandler,
	"signinGetHandler":        signinGetHandler,
	"signinSuccessGetHandler": signinSuccessGetHandler,
	"signinFatalGetHandler":   signinFatalGetHandler,
}

// Функция получает обработчик функции по имени
func getHandlerByName(handlerName string) gin.HandlerFunc {
	if res, ok := handlers[handlerName]; ok {
		return res
	}

	log.Printf("Неизвестный обработчик: %s", handlerName)
	return nil
}

func loginGetHandler(ctx *gin.Context) {
	body, err := os.ReadFile("cmd/pages/login-page/login.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

func rootPageGetHandler(ctx *gin.Context) {
	body, err := os.ReadFile("cmd/pages/main-page/main.html")
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
	body, err := os.ReadFile("cmd/pages/login-page/success.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	log.Printf("Handle post request for /login\n [username]: %s\t [password]: %s\n", username, password)
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninPostResponse struct {
	RedirectURL string `json:"redirectURL"`
}

func signinPostHandler(ctx *gin.Context) {
	// Чтение JSON данных из запроса
	var user User
	if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		log.Printf("unable to decode JSON: %v", err)
		// Обработка ошибки
		return
	}

	// Создание объекта Response с полем redirectURL
	response := SigninPostResponse{
		RedirectURL: "/signin-success",
	}

	log.Printf("Handle post request for /signin\n [username]: %s\t [email]: %s\t [password]: %s\n", user.Username, user.Email, user.Password)
	if err := serviceuser.RegisterUser(serviceuser.NewUser(user.Username, user.Email, user.Password)); err != nil {
		fmt.Print(err)
		log.Printf("%v", err)
		// TODO: в случае внутренней ошибки сервиса обеспечить редирект на страницу с ошибкой
		response = SigninPostResponse{
			RedirectURL: "/signin-fatal",
		}
	}

	// Отправка JSON-ответа с полем redirectURL
	ctx.JSON(http.StatusOK, response)
}
func signinGetHandler(ctx *gin.Context) {
	body, err := os.ReadFile("cmd/pages/signin-pages/signin.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

func signinSuccessGetHandler(ctx *gin.Context) {
	body, err := os.ReadFile("cmd/pages/signin-pages/signin-success.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

func signinFatalGetHandler(ctx *gin.Context) {
	body, err := os.ReadFile("cmd/pages/signin-pages/signin-fatal.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}
