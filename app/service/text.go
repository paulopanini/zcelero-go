package service

import (
	"fmt"
	"zcelero/app/datastruct"
	"zcelero/app/model"
	"zcelero/app/repository"
)

type TextService interface {
	CreateTextData(data model.TextData) model.TextData
	GetTextData(id string) model.TextData
	GetAllData() []model.TextData
}

type textService struct {
	db repository.DBSimulator
}

func NewTextService(simulator repository.DBSimulator) TextService {
	return &textService{
		db: simulator,
	}
}

func (t *textService) CreateTextData(data model.TextData) model.TextData {
	fmt.Println("I'm here")
	textData := datastruct.TextData{
		Text:       data.Data,
		Encryption: data.Encryption,
		KeySize:    data.KeySize,
	}

	dbData := t.db.AddData(textData)

	returnData := model.TextData{
		ID:         dbData.Id,
		Data:       dbData.Text,
		Encryption: dbData.Encryption,
		KeySize:    dbData.KeySize,
	}
	return returnData
}

func (t *textService) GetTextData(id string) model.TextData {
	dbData := t.db.GetDataById(id)
	returnData := model.TextData{
		ID:         dbData.Id,
		Data:       dbData.Text,
		Encryption: dbData.Encryption,
		KeySize:    dbData.KeySize,
	}
	return returnData
}

func (t *textService) GetAllData() []model.TextData {
	var response []model.TextData
	dataFromDb := t.db.GetAllData()

	for _, data := range dataFromDb {
		response = append(response, model.TextData{
			ID:         data.Id,
			Data:       data.Text,
			Encryption: data.Encryption,
			KeySize:    data.KeySize,
		})
	}
	return response
}
