package tests

import (
	"fmt"
	"github.com/stretchr/testify/mock"
)

type MockInterceptorService struct{ mock.Mock }

func (m *MockInterceptorService) GetLocation(distances ...float32) (x, y float32){
	args := m.Called(distances)
	if len(m.ExpectedCalls) == 0 {
		return 0, 0
	}
	if args.Get(0) == nil {
		return -1,-1
	}
	return args.Get(0).(float32), args.Get(1).(float32)
}

func (m *MockInterceptorService) GetMessage(messages ...[]string) (msg string){
	args := m.Called(messages)
	if len(m.ExpectedCalls) == 0 {
		return fmt.Sprint("Not DEFINED expected calls. ")
	}
	if args.Get(0) == nil {
		return ""
	}
	return args.Get(0).(string)
}