package interfaces

type DecodeInterface interface {
	GetLocation(distances ...float32) (x, y float32)
	GetMessage(messages ...[]string) (msg string)
}
