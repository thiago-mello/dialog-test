package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/config/server"
)

func main() {
	port := config.GetInt("server.port")
	if port == 0 {
		log.Fatal("config value server.port is mandatory")
	}

	e := server.GetServer()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	log.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
