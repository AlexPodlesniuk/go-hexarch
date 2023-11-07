package ports

import (
	"net/http"
	"user-alerts/app"
	"user-alerts/app/query"

	"github.com/labstack/echo/v4"
)

type Api struct {
	e   *echo.Echo
	app app.Application
}

func NewApi(app app.Application) *Api {
	e := echo.New()
	api := &Api{e, app}

	// Register endpoints
	api.registerRoutes()

	return api
}

func (api *Api) Start() {
	api.e.Logger.Fatal(api.e.Start(":1324"))
}

func (api *Api) registerRoutes() {

	api.e.GET("/users/:id", api.FindUser)
}

func (api *Api) FindUser(c echo.Context) error {
	id := c.Param("id")
	findUser := api.app.Queries.FindUser
	user, err := findUser.Handle(c.Request().Context(), &query.FindUser{ID: id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving item")
	}

	return c.JSON(http.StatusOK, user)
}

type CreateUserDto struct {
	Name string
}
