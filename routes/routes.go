package routes

import (
	"crud_api_mvc/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	controllers.InitializeDB()

	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.GetItems)
	r.GET("/items/:id", controllers.GetItem)
	r.PUT("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)

	return r
}
