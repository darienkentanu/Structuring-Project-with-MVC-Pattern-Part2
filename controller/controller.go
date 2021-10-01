package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/darienkentanu/Structuring-Project-with-MVC-Pattern-Part2/model"
	"github.com/labstack/echo"
)

var db = model.DB

// get all users
func GetUsersController(c echo.Context) error {
	var books []model.Book
	if err := db.Find(&books).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, books)
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var book model.Book
	if err := db.First(&book, bookId).Error; err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if book.ID == 0 {
		return c.String(http.StatusNotFound, "book not found")
	}
	return c.JSON(http.StatusOK, book)
}

// create new user
func CreateUserController(c echo.Context) error {
	book := model.Book{}
	if err := c.Bind(&book); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if err := db.Save(&book).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, book)
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var book model.Book
	if err := db.First(&book, bookId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if book.ID == 0 {
		return c.String(http.StatusNotFound, "book not found")
	}
	if err := db.Delete(&book).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, book)
}

func UpdateUserController(c echo.Context) error {
	// your solution here
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var book model.Book
	if err := db.First(&book, bookId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if book.ID == 0 {
		return c.String(http.StatusNotFound, "book not found")
	}
	if err := c.Bind(&book); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if err := db.Save(&book).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, book)
}
