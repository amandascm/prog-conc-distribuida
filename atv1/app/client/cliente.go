package main

import (
	"fmt"
	"os"
	"strconv"
	collectorproxy "test/atv1/distribution/proxies/collector"
	namingproxy "test/atv1/services/naming/proxy"
	"test/shared"
	"time"
)

func main() {
	Cliente()
}

func Cliente() {

	ClientePerf()
	os.Exit(0)

	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	collector := collectorproxy.New(naming.Find("Collector"))

	// Chamada remota ao Collector
	collector.Log("log message")
}

func ClientePerf() {
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	collector := collectorproxy.New(naming.Find("Collector"))

	for i := 0; i < shared.StatisticSample; i++ {
		t1 := time.Now()
		for j := 0; j < shared.SampleSize; j++ {
			collector.Log("this is a log message " + strconv.FormatInt(time.Since(t1).Milliseconds(), 10))
		}
		fmt.Println(i, "sample;", time.Since(t1).Milliseconds())
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Experiemnt finalised...")
}
