package main

import (
	calculatorinvoker "test/atv1/distribution/invokers/calculatorinvoker"
	namingproxy "test/atv1/services/naming/proxy"
	"test/shared"
)

func main() {
	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Create instance of invokers
	calculatorInv := calculatorinvoker.New(shared.LocalHost, shared.CalculatorPort)

	// Register services in Naming
	naming.Bind("Calculator", shared.NewIOR(calculatorInv.Ior.Host, calculatorInv.Ior.Port))

	// Invoke services
	calculatorInv.Invoke()
}
