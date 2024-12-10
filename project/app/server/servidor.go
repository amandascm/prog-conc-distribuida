package main

import (
	"test/project/distribution/invokers/calculatorinvoker"
	lifecycle "test/project/distribution/lifecycle"
	"test/project/infrastructure/srh"
	namingproxy "test/project/services/naming/proxy"
	"test/shared"
)

func main() {
	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Creating the lifecycle manager
	lifecycleManager := lifecycle.NewLifecycleManager()
	defer lifecycleManager.Finish()

	// Create calculatorInvoker instance
	calculatorInvoker := calculatorinvoker.New(shared.LocalHost, shared.CalculatorPort, *lifecycleManager)

	// Create SRH instance
	srh := srh.NewWithInvoker(shared.LocalHost, shared.CalculatorPort, &calculatorInvoker)

	// Register services in Naming
	naming.Bind("Calculator", shared.NewIOR(calculatorInvoker.Ior.Host, calculatorInvoker.Ior.Port))

	srh.Serve()
}
