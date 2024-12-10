package lifecycle

import (
	"fmt"
	"test/atv1/app/businesses/calculator"
	management "test/atv1/distribution/lifecycle/management"
	"test/atv1/distribution/lifecycle/management/calculatorpool"
	"test/shared"
)

type LifecycleManager struct {
	strategy management.LifecycleManagementStrategy
}

func NewLifecycleManager() *LifecycleManager {
	var strategy management.LifecycleManagementStrategy
	switch shared.LifecycleManagementStrategy {
	case "pooling":
		pool := calculatorpool.NewObjectPool()
		strategy = pool.Activate()
	default:
		panic("Undefined management strategy")
	}
	return &LifecycleManager{strategy: strategy}
}

func (lm *LifecycleManager) GetObject() calculator.Calculator {
	return lm.strategy.Get()
}

func (lm *LifecycleManager) ReleaseObject(calc *calculator.Calculator) {
	lm.strategy.Release(calc)
}

func (lm *LifecycleManager) Finish() {
	fmt.Println("Destroying Lifecycle Manager and releasing all objects.")
	lm.strategy.Deactivate()
}
