package utils

import (
	"ejercicio/commons/models"
	"github.com/sirupsen/logrus"
)

func ToFloat32From64(value float64) float32{
	return float32(value)
}

func ToFloat32FromInt(value int) float32{
	return float32(value)
}

func GenerateListInformation(input []models.Satellite) ([]float32, [][]string){
	var listDistances []float32
	var listMessages [][]string

	for _, v := range input{
		listDistances = append(listDistances, v.Distance)
		listMessages = append(listMessages, v.Message)
	}
	logrus.Info("listDistances : ", listDistances)
	logrus.Info("listMessages : ", listMessages)

	return listDistances, listMessages
}