package repository

import (
	"strconv"
	datastruct "zcelero/app/datastruct"
)

type DBSimulator interface {
	AddData(data datastruct.TextData) datastruct.TextData
	GetDataById(id string) datastruct.TextData
	GetAllData() []datastruct.TextData
}

type dbSimulator struct{}

func NewDBSimulator() DBSimulator {
	return &dbSimulator{}
}

var tableValues []datastruct.TextData

func (d *dbSimulator) AddData(data datastruct.TextData) datastruct.TextData {
	id := len(tableValues) + 1
	data.Id = strconv.Itoa(id)
	tableValues = append(tableValues, data)
	return data
}

func (d *dbSimulator) GetDataById(id string) datastruct.TextData {
	for _, row := range tableValues {
		if row.Id == id {
			return row
		}
	}
	return datastruct.TextData{}
}

func (d *dbSimulator) GetAllData() []datastruct.TextData {
	return tableValues
}
