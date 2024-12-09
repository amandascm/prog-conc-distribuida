package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	calculatorproxy "test/atv1/distribution/proxies/calculatorproxy"
	namingproxy "test/atv1/services/naming/proxy"
	"test/shared"
	"time"
)

func main() {
	ClientePerf()
}

func Cliente() {

	ClientePerf()
	os.Exit(0)

	// Obtain proxies
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	calculator := calculatorproxy.New(naming.Find("Calculator"))

	// Chamada remota ao calculator
	calculator.Som(2, 2)
}

func worker(reqsChan chan [2]int, calculator calculatorproxy.CalculatorProxy, wg *sync.WaitGroup) {
	for args := range reqsChan {
		wg.Add(1)
		log.Println("Sending request with", args[0], args[1])
		response := calculator.Som(args[0], args[1])
		log.Println("Received response for request", args[0], args[1], response)
		wg.Done()
	}
}

func ClientePerf() {
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	calculator := calculatorproxy.New(naming.Find("Calculator"))

	lCh := make(chan [2]int)
	wg := new(sync.WaitGroup)

	// We always batch send twice the pool size of requests
	for i := 0; i < shared.PoolSize*2; i++ {
		go worker(lCh, calculator, wg)
	}

	for i := 0; i < shared.StatisticSample; i++ {
		t1 := time.Now()
		for j := 0; j < shared.SampleSize; j++ {
			lCh <- [2]int{i, j}
		}
		wg.Wait()

		fmt.Println(i, "sample;", time.Since(t1).Milliseconds())
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Experiemnt finalised...")
}