package managers

import (
	"github.com/labstack/echo"
	"github.com/jgroeneveld/losmentor/api"
)

func ListManagers(ctx echo.Context) error {
	return ctx.JSON(200, api.PagedCollectionResponse{
		Elements: []*Manager{{ID: 12, FirstName: "Jon", LastName: "Snow"}},
		Page:     fakePageInfo(),
	})
}

func GetManager(ctx echo.Context) error {
	return ctx.JSON(200, &Manager{ID: 12, FirstName: "Jon", LastName: "Snow"})
}

func fakePageInfo() api.PageInfo {
	return api.PageInfo{
		PageNumber:    0,
		IsLastPage:    true,
		TotalElements: 1,
		TotalPages:    1,
		PageSize:      20,
	}
}
