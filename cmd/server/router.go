package server

/*
 * Project: I-wish-you
 * Created Date: Tuesday, July 11th 2023, 12:47:25 am
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
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func newRouter() *Router {
	return &Router{
		router: gin.Default(),
	}
}

func (r *Router) configure() {
	// Открываем файл конфигурации
	configFile, err := os.Open("configs/routes_config.json")
	if err != nil {
		log.Fatal("Не удалось открыть файл конфигурации: ", err)
	}
	defer configFile.Close()

	// Декодируем содержимое файла конфигурации в структуру
	var routesConfig []Route
	if err = json.NewDecoder(configFile).Decode(&routesConfig); err != nil {
		log.Fatal("Ошибка при декодировании файла конфигурации: ", err)
	}

	// Обрабатываем каждый маршрут из файла конфигурации
	for _, route := range routesConfig {
		// Регистрируем маршрут и обработчик в маршрутизаторе
		registerHandler(r.router, route)
	}
}

func (r *Router) run(addr string) error {
	return r.router.Run(addr)
}

func registerHandler(router *gin.Engine, route Route) {
	// Получаем соответствующий обработчик функции по имени
	handlerFunc := getHandlerByName(route.Handler)
	switch route.Method {
	case "GET":
		if route.Access == "private" {
			router.GET(route.Path, handlerFunc, authorizeMiddleware)
		} else {
			router.GET(route.Path, handlerFunc)
		}
	case "POST":
		if route.Access == "private" {
			router.POST(route.Path, handlerFunc, authorizeMiddleware)
		} else {
			router.POST(route.Path, handlerFunc)
		}
	default:
		log.Printf("Неизвестный метод: %s", route.Method)
	}
}

// TODO: выделить в интерфейс и сделать фабрику в зависимости от первого слова заголовка Authorization
func authorizeMiddleware(ctx *gin.Context) {
	log.Println("authorize Middleware is called")
	authHeader := ctx.GetHeader("Authorization")
	log.Printf("Header %s", authHeader)

	if authHeader == "" {
		ctx.String(401, "Требуется авторизация")
		ctx.Abort()
		return
	}
	// TODO: реализовать обработку ошибок на этапе авторизации
	// TODO: Придумать как отправлять запрос на редирект на главную страницу
	// TODO: Реализовать нормальный запрос с закодированными данными в заголовке на javascript

	// Проверяем, что заголовок Authorization начинается с "Basic "
	if strings.HasPrefix(authHeader, "Basic ") {
		// Извлекаем токен из строки "Basic base64(username:password)"
		token := strings.TrimPrefix(authHeader, "Basic ")

		// Декодируем base64
		decodedToken, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			ctx.String(400, "Ошибка при декодировании токена")
			ctx.Abort()
			return
		}

		// Преобразуем в строку и разделяем имя пользователя и пароль
		credentials := strings.SplitN(string(decodedToken), ":", 2)
		if len(credentials) != 2 {
			ctx.String(400, "Ошибка при извлечении имени пользователя и пароля")
			ctx.Abort()
			return
		}

		email := credentials[0]
		password := credentials[1]
		log.Printf("email %s and password %s from headers\n", email, password)
		//		if ok, _ := service.AuthorizeUser(service.ServiceUser{Name: "", Email: email, Password: password}); ok {
		ctx.Next()
		// } else {
		// 	ctx.Abort()
		// }

	} else {
		ctx.String(401, "Неподдерживаемая схема аутентификации")
		ctx.Abort()
	}
}
