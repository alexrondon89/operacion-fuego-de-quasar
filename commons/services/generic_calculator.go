package services

import (
	utils2 "ejercicio/commons/utils"
	"github.com/sirupsen/logrus"
	"math"
)

type GenericCalculator struct {

}

func (gc *GenericCalculator) Coordinates(distance float32) map[string]float32{
	logrus.Info("calculating enemy position from distance ", distance)
	x, y := gc.distanceCalculator(distance)
	return map[string]float32{"x": utils2.ToFloat32From64(x), "y": utils2.ToFloat32From64(y)}
}

func (gc *GenericCalculator) GetMessageOrdered(messages [][]string) []string{
	list := gc.getListOfLists(messages)
	message := []string{}

	for _,v := range list{
		for _, n:= range v{
			if z:= gc.contains(message, n); !z{
				message = append(message, n)
			}
		}
	}

	return message
}

func (gc *GenericCalculator) getListOfLists(messages [][]string) [10][10]string{
	listOfLits := [10][10]string{{}}
	for i, v := range messages{
		for y, z := range v{
			listOfLits[y][i] = z
		}
	}
	return listOfLits
}

func (gc *GenericCalculator) contains(s []string, e string) bool {
	for _, a := range s {
		if a == e || e == "" {
			return true
		}
	}
	return false
}

func (gc *GenericCalculator) distanceCalculator(distance float32) (float64, float64){
	y := math.Sqrt((0.8*(float64(distance*distance)))-40000)
	x := 300 -(0.5*y)
	return x, y
}
