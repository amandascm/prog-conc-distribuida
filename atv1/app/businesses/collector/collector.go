package collector

import (
	"fmt"
	"math/rand"
	"test/shared"
	"time"
)

type Collector struct {
	ID int
}

func (Collector) Som(a, b float64) float64 {
	time.Sleep(time.Duration(shared.SumTime * time.Millisecond))
	return a + b
}

func (Collector) Log(message string) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	// Save log to some index
	fmt.Printf("    Received log with message: %s\n", message)
}

func (Collector) Metric(metric string) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	// Save metric to some index
	fmt.Printf("    Received metric with message: %s\n", metric)
}
