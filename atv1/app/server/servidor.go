package main

import (
	"test/atv1/distribution/invokers/calculatorinvoker"
	"test/atv1/distribution/lifecyclemanager"
	"test/atv1/infrastructure/srh"
	namingproxy "test/atv1/services/naming/proxy"
	"test/shared"
)

func main() {
	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Creating the lifecycle manager
	lifecycleManager := lifecyclemanager.NewLifecycleManager(shared.PoolSize)
	defer lifecycleManager.Destroy()
	
	// Create calculatorInvoker instance
	calculatorInvoker := calculatorinvoker.New(shared.LocalHost, shared.CalculatorPort, *lifecycleManager)

	// Create SRH instance
	srh := srh.NewWithInvoker(shared.LocalHost, shared.CalculatorPort, &calculatorInvoker)

	// Register services in Naming
	naming.Bind("Calculator", shared.NewIOR(calculatorInvoker.Ior.Host, calculatorInvoker.Ior.Port))

	srh.Serve()
}