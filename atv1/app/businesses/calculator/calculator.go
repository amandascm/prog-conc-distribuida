package calculator

import (
	"test/shared"
	"time"
)

type Calculator struct {
	ID int
}

func (Calculator) Som(a, b float64) float64 {
	time.Sleep(time.Duration(shared.SumTime * time.Millisecond))
	return a + b
}

func (Calculator) Dif(a, b float64) float64 {
	time.Sleep(time.Duration(shared.SumTime * time.Millisecond))
	return a - b
}

func (Calculator) Mul(a, b float64) float64 {
	time.Sleep(time.Duration(shared.SumTime * time.Millisecond))
	return a * b
}

func (Calculator) Div(a, b float64) float64 {
	time.Sleep(time.Duration(shared.SumTime * time.Millisecond))
	return a / b
}
