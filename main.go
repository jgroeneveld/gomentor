package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Manager struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
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

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	managers := e.Group("/managers")
	managers.GET("", listManagers)
	managers.GET("/:id", getManager)

	e.Logger.Fatal(e.Start(":8000"))
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

func fakePageInfo() PageInfo {
	return PageInfo{
		PageNumber:    0,
		IsLastPage:    true,
		TotalElements: 1,
		TotalPages:    1,
		PageSize:      20,
	}
}
