package collectorpool

import (
	"fmt"
	"test/atv1/app/businesses/collector"
)

type CollectorPool struct {
	pool chan *collector.Collector
}

func NewObjectPool(size int) *CollectorPool {
	pool := make(chan *collector.Collector, size)
	for i := range size {
		pool <- &collector.Collector{ID: i}
	}
	return &CollectorPool{
		pool: pool,
	}
}

func (calcPool *CollectorPool) Get() *collector.Collector {
	obj := <-calcPool.pool
	fmt.Printf("Got object available at channel with ID: %d\n", obj.ID)
	return obj
}

func (calcPool *CollectorPool) Put(obj *collector.Collector) {
	fmt.Printf("Returned to pool object with ID: %d\n", obj.ID)
	calcPool.pool <- obj
}
