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
	"os"

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
		router.GET(route.Path, handlerFunc)
	case "POST":
		router.POST(route.Path, handlerFunc)
	default:
		log.Printf("Неизвестный метод: %s", route.Method)
	}
}
