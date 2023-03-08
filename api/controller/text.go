package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zcelero/app/model"
	"zcelero/app/service"
)

type TextDataController interface {
	CreateTextData(c *gin.Context)
	GetTextData(c *gin.Context)
	GetAllData(c *gin.Context)
}

type textDataController struct {
	service service.TextService
}

func NewTextDataController(textService service.TextService) TextDataController {
	return &textDataController{
		service: textService,
	}
}

func (t *textDataController) CreateTextData(c *gin.Context) {

	var newTextData model.TextData
	if err := c.BindJSON(&newTextData); err != nil {
		return
	}
	data := t.service.CreateTextData(newTextData)
	c.IndentedJSON(http.StatusCreated, data)
}

func (t *textDataController) GetTextData(c *gin.Context) {
	id := c.Param("id")

	data := t.service.GetTextData(id)

	c.IndentedJSON(http.StatusCreated, data)
}

func (t *textDataController) GetAllData(c *gin.Context) {
	data := t.service.GetAllData()

	c.IndentedJSON(http.StatusCreated, data)
}
