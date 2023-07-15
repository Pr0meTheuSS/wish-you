package server

// REFACTOR:  перенести полезную нагрузку работы хэндлеров на уровень сервисов
// REFACTOR:  перенести используемые типы в отдельный файл
// FIX: удалить log.Fatalf для случая невозможночти чтения файла страницы

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
	service "main/cmd/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var handlers = map[string]gin.HandlerFunc{
	"rootPageGetHandler":      rootPageGetHandler,
	"loginGetHandler":         loginGetHandler,
	"loginPostHandler":        loginPostHandler,
	"sendMessagePostHandler":  sendMessagePostHandler,
	"signupPostHandler":       signupPostHandler,
	"signupGetHandler":        signupGetHandler,
	"signupSuccessGetHandler": signupSuccessGetHandler,
	"signupFatalGetHandler":   signupFatalGetHandler,
	"welcomeGetHandler":       welcomeGetHandler,
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
	body, err := os.ReadFile("cmd/pages/login-pages/login.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

func rootPageGetHandler(ctx *gin.Context) {
	body, err := os.ReadFile("cmd/pages/main-pages/main.html")
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
	body, err := os.ReadFile("cmd/pages/login-pages/success.html")
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

type RedirectPostResponse struct {
	RedirectURL string `json:"redirectURL"`
}

func signupPostHandler(ctx *gin.Context) {
	// Чтение JSON данных из запроса
	var user User
	if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		log.Printf("unable to decode JSON: %v", err)
		// Обработка ошибки
		return
	}

	// Создание объекта Response с полем redirectURL
	response := RedirectPostResponse{
		RedirectURL: "/signup-success",
	}

	log.Printf("Handle post request for /signup\n [username]: %s\t [email]: %s\t [password]: %s\n", user.Username, user.Email, user.Password)
	if err := service.RegisterUser(service.NewUser(user.Username, user.Email, user.Password)); err != nil {
		fmt.Print(err)
		log.Printf("%v", err)
		// TODO: в случае внутренней ошибки сервиса обеспечить редирект на страницу с ошибкой
		response = RedirectPostResponse{
			RedirectURL: "/signup-fatal",
		}
	}

	// Отправка JSON-ответа с полем redirectURL
	ctx.JSON(http.StatusOK, response)
}
func signupGetHandler(ctx *gin.Context) {
	body, err := os.ReadFile("cmd/pages/signup-pages/signup.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

func signupSuccessGetHandler(ctx *gin.Context) {
	body, err := os.ReadFile("cmd/pages/signup-pages/signup-success.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

func signupFatalGetHandler(ctx *gin.Context) {
	body, err := os.ReadFile("cmd/pages/signup-pages/signup-fatal.html")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}

func welcomeGetHandler(ctx *gin.Context) {
	body, err := service.LoadPage("welcome-page")
	if err != nil {
		log.Printf("Page loading finished wtih error %v", err)
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte{})
		return
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}
