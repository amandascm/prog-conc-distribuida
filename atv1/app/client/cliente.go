package main

import (
	"fmt"
	"log"
	"os"
	"sync"
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

func worker(linkChan chan int, collector collectorproxy.CollectorProxy, wg *sync.WaitGroup) {
	defer wg.Done()
	for args := range linkChan {
		collector.Som(0, args)
		log.Println("Received response for request ", args)
	}
 }
 
func ClientePerf() {
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	collector := collectorproxy.New(naming.Find("Collector"))

    lCh := make(chan int)
    wg := new(sync.WaitGroup)

    // Adding routines to workgroup and running then
    for i := 0; i < 30; i++ {
        wg.Add(1)
        go worker(lCh, collector, wg)
    }

	for i := 0; i < shared.StatisticSample; i++ {
		t1 := time.Now()
		for j := 0; j < shared.SampleSize; j++ {
			lCh <- j
		}
		fmt.Println(i, "sample;", time.Since(t1).Milliseconds())

		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Experiemnt finalised...")
}
