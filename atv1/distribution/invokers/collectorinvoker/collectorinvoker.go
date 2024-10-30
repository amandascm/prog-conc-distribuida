package collectorinvoker

import (
	"log"
	collectorpool "test/atv1/distribution/collectorpool"
	"test/atv1/distribution/marshaller"
	"test/atv1/distribution/miop"
	"test/atv1/infrastructure/srh"
	"test/shared"
)

type CollectorInvoker struct {
	Ior shared.IOR
}

func New(h string, p int) CollectorInvoker {
	ior := shared.IOR{Host: h, Port: p}
	inv := CollectorInvoker{Ior: ior}
	// Instantiate pool

	return inv
}

func (i CollectorInvoker) Invoke() {
	s := srh.NewSRH(i.Ior.Host, i.Ior.Port)
	m := marshaller.Marshaller{}
	miopPacket := miop.Packet{}

	// Create a pool of instances of Collector
	pool := collectorpool.NewObjectPool(10)

	for {
		// Invoke SRH
		b := s.Receive()

		// Unmarshall miop packet
		miopPacket = m.Unmarshall(b)

		// Extract request from publisher
		r := miop.ExtractRequest(miopPacket)

		_p1 := string(r.Params[0].(string))

		// Get instance from pool
		c := pool.Get()

		switch r.Op {
		case "Log":
			go func() {
				// Release instance (put back in pool)
				defer pool.Put(c)
				c.Log(_p1)
			}()
		case "Metric":
			go func() {
				// Release instance (put back in pool)
				defer pool.Put(c)
				c.Metric(_p1)
			}()
		default:
			log.Fatal("Invoker:: Operation '" + r.Op + "' is unknown:: ")
		}

		// Prepare reply
		var params []interface{}

		// Create miop reply packet
		miop := miop.CreateReplyMIOP(params)

		// Marshall miop packet
		b = m.Marshall(miop)

		// Send marshalled packet
		s.Send(b)
	}
}
