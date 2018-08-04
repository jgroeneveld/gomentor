package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Manager struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PagedCollectionResponse struct {
	Elements interface{} `json:"elements"`
	Page     PageInfo    `json:"page"`
}

type PageInfo struct {
	PageNumber    int  `json:"page_number"`
	IsLastPage    bool `json:"is_last_page"`
	TotalElements int  `json:"total_elements"`
	TotalPages    int  `json:"total_pages"`
	PageSize      int  `json:"page_size"`
}

type QuoteApiResponse struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	api := e.Group("/api")

	api.GET("/random_quote", randomQuote)

	managers := api.Group("/managers")
	managers.GET("", listManagers)
	managers.GET("/:id", getManager)

	e.Logger.Fatal(e.Start(":" + port))
}

func listManagers(ctx echo.Context) error {
	return ctx.JSON(200, PagedCollectionResponse{
		Elements: []*Manager{{ID: 12, FirstName: "Jon", LastName: "Snow"}},
		Page:     fakePageInfo(),
	})
}

func getManager(ctx echo.Context) error {
	return ctx.JSON(200, &Manager{ID: 12, FirstName: "Jon", LastName: "Snow"})
}

func randomQuote(ctx echo.Context) error {
	url := "https://api.forismatic.com/api/1.0/?method=getQuote&key=457635&format=json&lang=en"

	quoteResponse := QuoteApiResponse{}
	err := getJSON(url, &quoteResponse)
	if err != nil {
		return err
	}

	quote := Quote{
		Author: quoteResponse.QuoteAuthor,
		Text:   quoteResponse.QuoteText,
	}

	return ctx.JSON(200, quote)
}

func getJSON(url string, target interface{}) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}

func fakePageInfo() PageInfo {
	return PageInfo{
		PageNumber:    0,
		IsLastPage:    true,
		TotalElements: 1,
		TotalPages:    1,
		PageSize:      20,
	}
}
