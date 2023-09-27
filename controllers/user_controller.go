package controllers

import (
	"crud_api_mvc/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbuser *gorm.DB

func InitializeUserDB() {
	var err error
	dbuser, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang_crud password=root sslmode=disable")
	if err != nil {
		panic("Failed to connect to database")
	}
	dbuser.AutoMigrate(&models.User{})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbuser.Create(&user)
	c.JSON(http.StatusCreated, user)
}
