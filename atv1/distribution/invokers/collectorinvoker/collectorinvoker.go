package collectorinvoker

import (
	"log"
	"net"
	collectorpool "test/atv1/distribution/collectorpool"
	"test/atv1/distribution/marshaller"
	"test/atv1/distribution/miop"
	"test/atv1/infrastructure/srh"
	"test/shared"
	"time"
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
	pool := collectorpool.NewObjectPool(shared.PoolSize)

	for {
		// Invoke SRH
		b, conn := s.Receive()

		go func (conn net.Conn) {
			// Unmarshall miop packet
			miopPacket = m.Unmarshall(b)
	
			// Extract request from publisher
			r := miop.ExtractRequest(miopPacket)
	
			_p1 := float64(r.Params[0].(float64))
			_p2 := float64(r.Params[1].(float64))
	
			// Get instance from pool
			c := pool.Get()
	
			// Prepare reply
			var params []interface{}
	
			switch r.Op {
			case "Som":
				params = append(params, _p1 + _p2)
				time.Sleep(time.Duration(shared.SumTime * time.Millisecond))
				pool.Put(c)
			default:
				log.Fatal("Invoker:: Operation '" + r.Op + "' is unknown:: ")
			}
	
			// Create miop reply packet
			miop := miop.CreateReplyMIOP(params)
	
			// Marshall miop packet
			b = m.Marshall(miop)
	
			// Send marshalled packet
			s.Send(b, conn)
		}(conn)
	}
}
