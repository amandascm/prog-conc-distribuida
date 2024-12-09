package main

import (
	"test/atv1/infrastructure/srh"
	namingproxy "test/atv1/services/naming/proxy"
	"test/shared"
)

func main() {
	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Create SRH instance
	srh := srh.NewSRH(shared.LocalHost, shared.CalculatorPort)

	// Register services in Naming
	naming.Bind("Calculator", shared.NewIOR(srh.Invoker.Ior.Host, srh.Invoker.Ior.Port))

	srh.Serve()
}
