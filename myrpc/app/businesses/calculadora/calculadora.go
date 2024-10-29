package calculadora

import (
	"math/rand"
	"time"
)

type Calculadora struct {
	ID int
}

func (Calculadora) Som(p1, p2 int) int {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return p1 + p2
}

func (Calculadora) Dif(p1, p2 int) int {
	return p1 - p2
}

func (Calculadora) Mul(p1, p2 int) int {
	return p1 * p2
}

func (Calculadora) Div(p1, p2 int) int {
	if p2 == 0 {
		return 0
	} else {
		return p1 / p2
	}
}
