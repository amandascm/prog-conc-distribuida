package main

import (
	collectorinvoker "test/atv1/distribution/invokers/collectorinvoker"
	namingproxy "test/atv1/services/naming/proxy"
	"test/shared"
)

func main() {
	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)

	// Create instance of invokers
	collectorInv := collectorinvoker.New(shared.LocalHost, shared.CollectorPort)

	// Register services in Naming
	naming.Bind("Collector", shared.NewIOR(collectorInv.Ior.Host, collectorInv.Ior.Port))

	// Invoke services
	collectorInv.Invoke()
}
