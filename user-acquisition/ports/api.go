package ports

import (
	"net/http"
	"user-acquisition/app"
	"user-acquisition/app/command"
	"user-acquisition/app/query"

	"github.com/google/uuid"
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
	api.e.Logger.Fatal(api.e.Start(":1323"))
}

func (api *Api) registerRoutes() {

	api.e.GET("/users/:id", api.FindUser)
	api.e.POST("/users", api.CreateUser)
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

func (api *Api) CreateUser(c echo.Context) error {
	var createUserDto CreateUserDto
	if err := c.Bind(&createUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid item data")
	}

	createUser := api.app.Commands.CreateUser
	createUserCmd := command.CreateUser{
		ID:   (uuid.New()).String(),
		Name: createUserDto.Name,
	}

	user, err := createUser.Handle(c.Request().Context(), &createUserCmd)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error retrieving item")
	}

	return c.JSON(http.StatusCreated, user)
}

type CreateUserDto struct {
	Name string
}
