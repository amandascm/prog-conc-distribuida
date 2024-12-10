package management

import "test/atv1/app/businesses/calculator"

type LifecycleManagementStrategy interface {
	Activate() LifecycleManagementStrategy
	Deactivate()
	Get() calculator.Calculator
	Release(obj *calculator.Calculator)
}
