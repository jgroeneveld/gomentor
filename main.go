package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"log"
	"github.com/jgroeneveld/losmentor/quotes"
	"github.com/jgroeneveld/losmentor/managers"
	"github.com/jgroeneveld/losmentor/fetching"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	dependencies := &dependencyConfiguration{
		jsonFetcher: fetching.NewJSONFetcher(),
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	configureHandlers(e, dependencies)

	e.Logger.Fatal(e.Start(":" + port))
}

func configureHandlers(e *echo.Echo, dependencies *dependencyConfiguration) {
	e.Static("/", "static")

	apiGroup := e.Group("/api")

	apiGroup.GET("/random_quote", (&quotes.RandomQuoteController{dependencies}).Handle)

	managersGroup := apiGroup.Group("/managers")
	managersGroup.GET("", managers.ListManagers)
	managersGroup.GET("/:id", managers.GetManager)
}

type dependencyConfiguration struct {
	jsonFetcher fetching.JSONFetcher
}

func (deps *dependencyConfiguration) JSONFetcher() fetching.JSONFetcher {
	if deps.jsonFetcher == nil {
		panic("JSONFetcher not configured")
	}
	return deps.jsonFetcher
}
