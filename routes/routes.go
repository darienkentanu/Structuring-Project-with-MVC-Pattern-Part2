package routes

import (
	c "github.com/darienkentanu/Structuring-Project-with-MVC-Pattern-Part2/controller"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// user routing
	e.GET("/books", c.GetUsersController)
	e.GET("/books/:id", c.GetUserController)
	e.POST("/books", c.CreateUserController)
	e.DELETE("/books/:id", c.DeleteUserController)
	e.PUT("/books/:id", c.UpdateUserController)

	return e
}
