package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/leandro-andrade-candido/api-go/docs"
	"github.com/leandro-andrade-candido/api-go/src/config"
	"github.com/leandro-andrade-candido/api-go/src/config/server"
	"github.com/leandro-andrade-candido/api-go/src/routes"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title		Simple Social Network API
// @version		1.0
// @description	This is an API for a very simple text based social network.
//
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				A Bearer token used for user authorization
func main() {
	port := config.GetInt("server.port")
	if port == 0 {
		log.Fatal("config value server.port is mandatory")
	}

	e := server.GetServer()
	routes.SetupRoutes(e)

	tp := server.GetTracerProvider()
	if tp != nil {
		defer func() {
			if err := tp.Shutdown(context.Background()); err != nil {
				log.Printf("Error shutting down tracer provider: %v", err)
			}
		}()
	}

	// Swagger setup
	e.GET("/docs/*", echoSwagger.WrapHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
