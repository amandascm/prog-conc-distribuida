package calculatorpool

import (
	"fmt"
	"test/atv1/app/businesses/calculator"
)

type CalculatorPool struct {
	pool chan *calculator.Calculator
}

func NewObjectPool(size int) *CalculatorPool {
	pool := make(chan *calculator.Calculator, size)
	for i := range size {
		pool <- &calculator.Calculator{ID: i}
	}
	return &CalculatorPool{
		pool: pool,
	}
}

func (calcPool *CalculatorPool) Get() *calculator.Calculator {
	obj := <-calcPool.pool
	fmt.Printf("Got object available at channel with ID: %d\n", obj.ID)
	return obj
}

func (calcPool *CalculatorPool) Put(obj *calculator.Calculator) {
	fmt.Printf("Returned to pool object with ID: %d\n", obj.ID)
	calcPool.pool <- obj
}
