package calculatorpool

import (
	"fmt"
	"test/atv1/app/businesses/calculator"
	management "test/atv1/distribution/lifecycle/management"
	"test/shared"
)

type CalculatorPool struct {
	Pool chan *calculator.Calculator
}

// Activate implements LifecycleManagementStrategy
func (calcPool *CalculatorPool) Activate() management.LifecycleManagementStrategy {
	return calcPool
}

// Deactivate implements LifecycleManagementStrategy
func (calcPool *CalculatorPool) Deactivate() {
	close(calcPool.Pool)
}

func NewObjectPool() *CalculatorPool {
	size := shared.PoolSize
	pool := make(chan *calculator.Calculator, size)
	for i := range size {
		pool <- &calculator.Calculator{ID: i}
	}
	return &CalculatorPool{
		Pool: pool,
	}
}

// Get implements LifecycleManagementStrategy
func (calcPool *CalculatorPool) Get() calculator.Calculator {
	obj := <-calcPool.Pool
	fmt.Printf("Got object available at channel with ID: %d\n", obj.ID)
	return *obj
}

// Release implements LifecycleManagementStrategy
func (calcPool *CalculatorPool) Release(obj *calculator.Calculator) {
	fmt.Printf("Returned to pool object with ID: %d\n", obj.ID)
	calcPool.Pool <- obj
}
