package main

import (
	"test/atv1/distribution/invokers/calculatorinvoker"
	"test/atv1/infrastructure/srh"
	namingproxy "test/atv1/services/naming/proxy"
	"test/shared"
)

func main() {
	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Create calculatorInvoker instance
	calculatorInvoker := calculatorinvoker.New(shared.LocalHost, shared.CalculatorPort)

	// Create SRH instance
	srh := srh.NewWithInvoker(shared.LocalHost, shared.CalculatorPort, &calculatorInvoker)

	// Register services in Naming
	naming.Bind("Calculator", shared.NewIOR(calculatorInvoker.Ior.Host, calculatorInvoker.Ior.Port))

	srh.Serve()
}
