package services

import (
	"ejercicio/commons/interfaces"
	"ejercicio/commons/utils"
	"github.com/sirupsen/logrus"
	"strings"
)

type GenericInterceptor struct {
	Calculator interfaces.CalculatorInterface
}

func (gi *GenericInterceptor) GetLocation(distances ...float32) (x, y float32){
	var coordinates map[string]float32
	var coordinatesList []map[string]float32

	for _, v := range distances{
		coordinates = gi.Calculator.Coordinates(v)
		coordinatesList = append(coordinatesList, coordinates)
	}

	logrus.Info("all coordinates: ", coordinatesList)
	cx, cy := utils.DistanceAverage(coordinatesList)
	logrus.Info("enemy coordinates ", "x: ",cx ,", y: ", cy)
	return cx, cy
}

func (gi *GenericInterceptor) GetMessage(messages ...[]string) (msg string){
	messageRecovered := gi.Calculator.GetMessageOrdered(messages)
	logrus.Info("enemy message: ", messageRecovered)
	return strings.Join(messageRecovered, " ")
}