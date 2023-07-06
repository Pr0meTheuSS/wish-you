package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: separate abstractions
// TODO: add graceful shutdown
// TODO: logging

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

func (a App) run() error {
	fmt.Printf(
		"Service %s is running, bro, server is listening port %d\n",
		a.configurations.serviceName,
		a.configurations.port)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, friends, application is running...")
	// })

	// http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, r.URL.Query().Get("message"))
	// })

	// http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	// 	body, err := ioutil.ReadFile("cmd/pages/login-page/login.html")
	// 	if err != nil {
	// 		log.Fatalf("unable to read file: %v", err)
	// 	}
	// 	fmt.Println(string(body))
	// 	fmt.Fprintf(w, string(body))
	// })

	// return http.ListenAndServe("127.0.0.1:6969", nil)

	router := gin.Default()

	router.GET("/login", func(ctx *gin.Context) {
		body, err := ioutil.ReadFile("cmd/pages/login-page/login.html")
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		fmt.Println(string(body))
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
	})

	router.POST("/login", func(ctx *gin.Context) {
		body, err := ioutil.ReadFile("cmd/pages/login-page/success.html")
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		// TODO: delete on release
		fmt.Printf("Handle post request for /login\n [username]: %s\t [password]: %s\n", username, password)

		ctx.Data(http.StatusOK, "text/html; charset=utf-8", body)
	})

	router.GET("/", func(ctx *gin.Context) {

	})

	return router.Run(":" + strconv.Itoa(a.configurations.port))
}

func loadConfigurations(configurations *Configurations) error {
	// TODO: implement - read from configs files
	*configurations = Configurations{
		port:        6969,
		serviceName: "wish-you",
	}

	return nil
}

func main() {
	appConfigurations := Configurations{}
	if loadConfigurations(&appConfigurations) != nil {
		// TODO: handle errors
	}

	wishYouApp := App{appConfigurations}
	if err := wishYouApp.run(); err != nil {
		// TODO: handle errors
		fmt.Print(err)
	}
}
