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
	"encoding/json"
	"log"
	internal "main/cmd/server/internal"
	service "main/cmd/services"
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

func getMethodRegister(router *gin.Engine, method string) func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	switch method {
	case "GET":
		return router.GET
	case "POST":
		return router.POST
	case "PATCH":
		return router.PATCH
	case "PUT":
		return router.PUT
	case "DELETE":
		return router.DELETE
	default:
		return nil
	}
}

func registerHandler(router *gin.Engine, route Route) {
	// Получаем соответствующий обработчик функции по имени
	handlerFunc := getHandlerByName(route.Handler)
	register := getMethodRegister(router, route.Method)
	if route.Access == "private" {
		register(route.Path, authorizeMiddleware, handlerFunc)
	} else if route.Access == "public" {
		register(route.Path, handlerFunc)
	}
}

// TODO: выделить в интерфейс и сделать фабрику в зависимости от первого слова заголовка Authorization
func authorizeMiddleware(ctx *gin.Context) {
	log.Println("authorize Middleware is called")
	// authHeader := ctx.GetHeader("Authorization")
	// log.Printf("Header %s", authHeader)

	authHeader, _ := ctx.Cookie("authToken")

	if authHeader == "" {
		// ctx.String(401, "Требуется авторизация")
		ctx.Redirect(302, "/login")
		ctx.Abort()
		return
	}

	// TODO: написать фабрику парсеров и обработчиков на каждый тип авторизации

	if strings.HasPrefix(authHeader, "Bearer ") {
		token := strings.TrimPrefix(authHeader, "Bearer ")
		// Валидация токена
		validToken, err := internal.ValidateToken(token)
		if err != nil {
			log.Println(err.Error())
			ctx.String(400, "Ошибка при декодировании токена")
			ctx.Abort()
			return
		}

		jwtClaims, err := internal.GetJwtClaims(validToken)
		if err != nil {
			log.Println(err.Error())
			ctx.String(400, "Ошибка при декодировании токена")
			ctx.Abort()
			return
		}

		if internal.IsJWTExpired(jwtClaims) {
			log.Println(err.Error())
			ctx.String(400, "Срок дейсвтия токена авторизации истекло")
			ctx.Abort()
			return

		}

		log.Printf("email %s and password %s from headers\n", jwtClaims.Email, jwtClaims.Password)
		if ok, _ := handler.UsersService.AuthenticateUser(service.AuthenticateUserRequest{
			Email:    jwtClaims.Email,
			Password: jwtClaims.Password,
		}); ok {
			ctx.Next()
		} else {
			ctx.Redirect(302, "/login")
			ctx.Abort()
		}

	} else {
		ctx.String(401, "Неподдерживаемая схема аутентификации")
		ctx.Abort()
	}
}
