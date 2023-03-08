package main

import (
	"zcelero/api"
	"zcelero/api/controller"
	"zcelero/app/repository"
	"zcelero/app/service"
)

func main() {

	db := repository.NewDBSimulator()
	textService := service.NewTextService(db)

	textController := controller.NewTextDataController(textService)

	api.Register(textController)
}
