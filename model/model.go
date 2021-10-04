package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	InitDB()
	InitialMigration()
}

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{"root", "password", "3306", "localhost", "crud_go"}
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username, config.DB_Password, config.DB_Host, config.DB_Port, config.DB_Name,
	)
	var err error
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&Book{})
}

type Book struct {
	gorm.Model
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}

type BooksModel struct {
	data []Book
}

func NewUserModel() *BooksModel {
	return &BooksModel{data: []Book{}}
}

func (bm *BooksModel) GetUsersController() error {
	return nil
}

func (bm *BooksModel) GetUserController() error {
	return nil
}

func (bm *BooksModel) CreateUserController() error {
	return nil
}

func (bm *BooksModel) DeleteUserController() error {
	return nil
}

func (bm *BooksModel) UpdateUserController() error {
	return nil
}
