package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
	calculatorproxy "test/project/distribution/proxies/calculatorproxy"
	namingproxy "test/project/services/naming/proxy"
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

func worker(reqsChan chan [2]int, calculator calculatorproxy.CalculatorProxy, wg *sync.WaitGroup, writer *csv.Writer) {
	for args := range reqsChan {
		wg.Add(1)
		log.Println("Sending request with", args[0], args[1])
		t1 := time.Now()
		response := calculator.Som(1234, 5678)
		elapsedTime := time.Since(t1).Nanoseconds()
		log.Println("Received response for request with args and latency: ", args[0], args[1], response, elapsedTime)
		// Write the elapsed time to the CSV file
		writer.Write([]string{
			fmt.Sprintf("%d", args[0]),     // Sample index
			fmt.Sprintf("%d", args[1]),     // Request index
			fmt.Sprintf("%d", response),    // Resulting integer
			fmt.Sprintf("%d", elapsedTime), // Elapsed time in milliseconds
		})
		wg.Done()
	}
}

func ClientePerf() {
	naming := namingproxy.New(shared.LocalHost, shared.NamingPort)
	calculator := calculatorproxy.New(naming.Find("Calculator"))

	lCh := make(chan [2]int)
	wg := new(sync.WaitGroup)

	// Open a CSV file to write the results
	file, err := os.Create(fmt.Sprintf("data/output/elapsed_times_pool%d.csv", shared.PoolSize))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	// Create a CSV writer
	writer := csv.NewWriter(file)
	// Write the header row to the CSV file
	writer.Write([]string{"Sample", "Request", "Response", "ElapsedTimeNs"})

	// We always batch send twice the pool size of requests
	for i := 0; i < shared.SampleSize; i++ {
		go worker(lCh, calculator, wg, writer)
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

	// Flush the CSV writer to ensure data is written
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error writing to CSV:", err)
	}

	fmt.Println("Experiemnt finalised...")
}
