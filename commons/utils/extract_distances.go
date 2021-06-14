package utils

func DistanceAverage(coordinates []map[string]float32) (float32, float32){
	var x float32 = 0.0
	var y float32 = 0.0

	for _, m := range coordinates{
		x += m["x"]
		y += m["y"]
	}
	long := ToFloat32FromInt(len(coordinates))
	x = x/long
	y = y/long

	return x, y
}


