package service

import (
	"zcelero/app/datastruct"
	"zcelero/app/model"
	"zcelero/app/repository"
)

type TextService interface {
	CreateTextData(data model.TextData) model.TextData
	GetTextData(id string, privateKey string) model.TextData
	GetAllData() []model.TextData
}

type textService struct {
	db         repository.DBSimulator
	encryption EncryptionService
}

func NewTextService(simulator repository.DBSimulator, service EncryptionService) TextService {
	return &textService{
		db:         simulator,
		encryption: service,
	}
}

func (t *textService) CreateTextData(data model.TextData) model.TextData {
	stringData, privateKey := t.getEncryptedData(data)
	textData := datastruct.TextData{
		Text:        stringData,
		IsEncrypted: data.ShouldEncrypt,
		KeySize:     data.KeySize,
	}

	dbData := t.db.AddData(textData)

	returnData := model.TextData{
		ID:         dbData.Id,
		PrivateKey: privateKey,
	}
	return returnData
}

func (t *textService) GetTextData(id string, privateKey string) model.TextData {
	dbData := t.db.GetDataById(id)

	decryptedData := t.getDecryptedData(dbData, privateKey)

	returnData := model.TextData{
		ID:            dbData.Id,
		Data:          decryptedData,
		ShouldEncrypt: dbData.IsEncrypted,
		KeySize:       dbData.KeySize,
	}
	return returnData
}

func (t *textService) GetAllData() []model.TextData {
	var response []model.TextData
	dataFromDb := t.db.GetAllData()

	for _, data := range dataFromDb {
		response = append(response, model.TextData{
			ID:            data.Id,
			Data:          data.Text,
			ShouldEncrypt: data.IsEncrypted,
			KeySize:       data.KeySize,
		})
	}
	return response
}

func (t *textService) getEncryptedData(data model.TextData) (string, string) {
	if data.ShouldEncrypt {
		return t.encryption.EncryptData(data.Data, data.KeySize)
	}
	return data.Data, ""
}

func (t *textService) getDecryptedData(data datastruct.TextData, privateKey string) string {
	if data.IsEncrypted {
		return t.encryption.DecryptData(data.Text, privateKey)
	}
	return data.Text
}
