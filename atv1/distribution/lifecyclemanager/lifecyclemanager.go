package lifecyclemanager

import (
	"fmt"
	"test/atv1/app/businesses/calculator"
	"test/atv1/distribution/calculatorpool"
)

type LifecycleManager struct {
	pool *calculatorpool.CalculatorPool
}

func NewLifecycleManager(maxPoolSize int) *LifecycleManager {
	pool := calculatorpool.NewObjectPool(10)
	return &LifecycleManager{pool: pool}
}

func (lm *LifecycleManager) GetObject() *calculator.Calculator {
	return lm.pool.Get()
}

func (lm *LifecycleManager) ReleaseObject(calc *calculator.Calculator) {
	lm.pool.Put(calc)
}

func (lm *LifecycleManager) Destroy() {
	fmt.Println("Destroying Lifecycle Manager and releasing all objects.")
	close(lm.pool.Pool)
}
