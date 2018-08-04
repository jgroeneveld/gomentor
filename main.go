package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"log"
	"github.com/jgroeneveld/losmentor/quotes"
	"github.com/jgroeneveld/losmentor/managers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	configureHandlers(e)

	e.Logger.Fatal(e.Start(":" + port))
}

func configureHandlers(e *echo.Echo) {
	e.Static("/", "static")

	apiGroup := e.Group("/api")

	apiGroup.GET("/random_quote", quotes.RandomQuote)

	managersGroup := apiGroup.Group("/managers")
	managersGroup.GET("", managers.ListManagers)
	managersGroup.GET("/:id", managers.GetManager)
}
