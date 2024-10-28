package calculadorapool

import (
	"fmt"
	"sync"
	"test/myrpc/app/businesses/calculadora"
)

type CalculadoraPool struct {
	pool   chan *calculadora.Calculadora
	mu     sync.Mutex
	nextID int
}

func NewObjectPool(size int) *CalculadoraPool {
	return &CalculadoraPool{
		pool: make(chan *calculadora.Calculadora, size),
	}
}

func (calcPool *CalculadoraPool) Get() *calculadora.Calculadora {
	calcPool.mu.Lock()
	defer calcPool.mu.Unlock()
	select {
	case obj := <-calcPool.pool:
		fmt.Printf("Got object available at channel with ID: %d\n", obj.ID)
		return obj
	default:
		calcPool.nextID++
		fmt.Printf("Got new object with ID: %d\n", calcPool.nextID)
		return &calculadora.Calculadora{ID: calcPool.nextID}
	}
}

func (calcPool *CalculadoraPool) Put(obj *calculadora.Calculadora) {
	select {
	case calcPool.pool <- obj:
	default:
		// No need if pool is full
	}
}
