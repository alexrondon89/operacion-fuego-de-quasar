package interfaces

type CalculatorInterface interface {
	Coordinates(distance float32) map[string]float32
	GetMessageOrdered(messages [][]string) []string
}

