package routes

import (
	c "github.com/darienkentanu/Structuring-Project-with-MVC-Pattern-Part2/controller"
	m "github.com/darienkentanu/Structuring-Project-with-MVC-Pattern-Part2/model"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// user routing
	mdl := m.NewUserModel()
	bc := c.NewController(mdl)

	e.GET("/books", bc.GetUsersController)
	e.GET("/books/:id", bc.GetUserController)
	e.POST("/books", bc.CreateUserController)
	e.DELETE("/books/:id", bc.DeleteUserController)
	e.PUT("/books/:id", bc.UpdateUserController)

	return e
}
