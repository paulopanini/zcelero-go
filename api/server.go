package api

import (
	"github.com/gin-gonic/gin"
	"zcelero/api/controller"
)

func Register(dataController controller.TextDataController) {

	router := gin.Default()
	router.POST("/text-management", dataController.CreateTextData)
	router.GET("/text-management/:id", dataController.GetTextData)
	router.GET("/text-management", dataController.GetAllData)

	router.Run("localhost:8080")

}
