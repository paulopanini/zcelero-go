package main

import (
	"zcelero/api"
	"zcelero/api/controller"
	"zcelero/app/repository"
	"zcelero/app/service"
)

func main() {

	db := repository.NewDBSimulator()
	encryptionService := service.NewEncryptionService()
	textService := service.NewTextService(db, encryptionService)

	textController := controller.NewTextDataController(textService)

	api.Register(textController)
}
