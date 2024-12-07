package calculatorinvoker

import (
	"log"
	"net"
	calculatorpool "test/atv1/distribution/calculatorpool"
	"test/atv1/distribution/marshaller"
	"test/atv1/distribution/miop"
	"test/atv1/infrastructure/srh"
	"test/shared"
)

type CalculatorInvoker struct {
	Ior shared.IOR
	pool calculatorpool.CalculatorPool
}

func New(h string, p int) CalculatorInvoker {
	ior := shared.IOR{Host: h, Port: p}
	pool := calculatorpool.NewObjectPool(shared.PoolSize)
	inv := CalculatorInvoker{Ior: ior, pool: *pool}
	return inv
}

func (i CalculatorInvoker) Invoke() {
	s := srh.NewSRH(i.Ior.Host, i.Ior.Port)
	m := marshaller.Marshaller{}
	miopPacket := miop.Packet{}

	for {
		// Invoke SRH
		b, conn := s.Receive()

		go func(conn net.Conn) {
			// Unmarshall miop packet
			miopPacket = m.Unmarshall(b)

			// Extract request from publisher
			r := miop.ExtractRequest(miopPacket)

			_p1 := float64(r.Params[0].(float64))
			_p2 := float64(r.Params[1].(float64))

			// Get instance from pool
			c := i.pool.Get()

			// Prepare reply
			var params []interface{}

			switch r.Op {
			case "Som":
				params = append(params, c.Som(_p1, _p2))
				i.pool.Put(c)
			default:
				log.Fatal("Invoker:: Operation '" + r.Op + "' is unknown:: ")
			}

			// Create miop reply packet
			miop := miop.CreateReplyMIOP(params)

			// Marshall miop packet
			b = m.Marshall(miop)

			// Send marshalled packet
			s.Send(conn, b)
		}(conn)
	}
}
