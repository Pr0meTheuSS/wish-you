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
	"database/sql"
	"encoding/json"
	"log"
	repository "main/cmd/repositories"
	"main/cmd/server/hash"
	internal "main/cmd/server/internal"
	service "main/cmd/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

type Handler struct {
	UsersService service.UsersService
}

var handler = Handler{
	UsersService: &service.RealUsersService{
		UsersRepo: &repository.SqlxUsersRepository{
			DB: &sqlx.DB{
				DB:     &sql.DB{},
				Mapper: &reflectx.Mapper{},
			},
		},
	},
}

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
	"faviconGetHandler":       faviconGetHandler,
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
	returnTargetPageOrNotFound("login", ctx)
}

func rootPageGetHandler(ctx *gin.Context) {
	returnTargetPageOrNotFound("main", ctx)
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

	// TODO: Логика обработки сообщения

	// Ответ сервера
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func loginPostHandler(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	ok, err := handler.UsersService.AuthenticateUser(service.AuthenticateUserRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Println(err.Error())
		// TODO: implement returnInternalError(ctx)
		//		returnTargetPageOrNotFound("login-fail", ctx)
		return
	}

	if !ok {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := internal.GenerateToken(email, password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		ctx.Abort()
		return
	}

	ctx.JSON(200, LoginResponse{
		Token: token,
	})
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RedirectPostResponse struct {
	RedirectURL string `json:"redirectURL"`
}

type LoginResponse struct {
	RedirectURL string `json:"redirectURL"`
	Token       string `json:"token"`
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

	if err := handler.UsersService.RegisterUser(service.NewUser(user.Username, user.Email, hash.GetHash(user.Password))); err != nil {
		log.Printf("%v", err)
		response = RedirectPostResponse{
			RedirectURL: "/signup-fatal",
		}
	}

	// Отправка JSON-ответа с полем redirectURL
	ctx.JSON(http.StatusOK, response)
}

func signupGetHandler(ctx *gin.Context) {
	returnTargetPageOrNotFound("signup", ctx)
}

func signupSuccessGetHandler(ctx *gin.Context) {
	returnTargetPageOrNotFound("signup-success", ctx)
}

func signupFatalGetHandler(ctx *gin.Context) {
	returnTargetPageOrNotFound("signup-fatal", ctx)
}

func welcomeGetHandler(ctx *gin.Context) {
	returnTargetPageOrNotFound("welcome-page", ctx)
}

func pageNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Страница не найдена"})
}

func faviconGetHandler(ctx *gin.Context) {
	ctx.File("favicon.ico")
}

func returnTargetPageOrNotFound(pageName string, ctx *gin.Context) {
	body, err := service.LoadPage(pageName)
	if err != nil {
		log.Printf("Page loading finished wtih error %v", err)
		pageNotFound(ctx)
		return
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
}
