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

	// Create invoker instance
	invoker := calculatorinvoker.New(shared.LocalHost, shared.CalculatorPort)

	// Create SRH instance
	srh := srh.NewWithInvoker(shared.LocalHost, shared.CalculatorPort, &invoker)

	// Register services in Naming
	naming.Bind("Calculator", shared.NewIOR(invoker.Ior.Host, invoker.Ior.Port))

	srh.Serve()
}
