package management

import "test/project/app/businesses/calculator"

type LifecycleManagementStrategy interface {
	Activate() LifecycleManagementStrategy
	Deactivate()
	Get() calculator.Calculator
	Release(obj *calculator.Calculator)
}
