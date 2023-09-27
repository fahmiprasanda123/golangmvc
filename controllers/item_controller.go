package controllers

import (
	"crud_api_mvc/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitializeDB() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang_crud password=root sslmode=disable")
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Item{})
}

func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&item)
	c.JSON(http.StatusCreated, item)
}

func GetItems(c *gin.Context) {
	var items []models.Item
	db.Find(&items)
	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var item models.Item
	db.First(&item, id)
	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func UpdateItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var item models.Item
	db.First(&item, id)
	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&item)
	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var item models.Item
	db.First(&item, id)
	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	db.Delete(&item)
	c.JSON(http.StatusNoContent, nil)
}
