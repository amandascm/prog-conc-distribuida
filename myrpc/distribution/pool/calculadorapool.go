package calculadorapool

import (
	"fmt"
	"test/myrpc/app/businesses/calculadora"
)

type CalculadoraPool struct {
	pool chan *calculadora.Calculadora
}

func NewObjectPool(size int) *CalculadoraPool {
	pool := make(chan *calculadora.Calculadora, size)
	for i := range size {
		pool <- &calculadora.Calculadora{ID: i}
	}
	return &CalculadoraPool{
		pool: pool,
	}
}

func (calcPool *CalculadoraPool) Get() *calculadora.Calculadora {
	obj := <-calcPool.pool
	fmt.Printf("Got object available at channel with ID: %d\n", obj.ID)
	return obj
}

func (calcPool *CalculadoraPool) Put(obj *calculadora.Calculadora) {
	fmt.Printf("Returned to pool object with ID: %d\n", obj.ID)
	calcPool.pool <- obj
}
