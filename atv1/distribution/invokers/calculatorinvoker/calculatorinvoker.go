package calculatorinvoker

import (
	"log"
	lifecyclemanager "test/atv1/distribution/lifecyclemanager"
	"test/atv1/distribution/marshaller"
	"test/atv1/distribution/miop"
	"test/shared"
)

type CalculatorInvoker struct {
	Ior shared.IOR
	lm  lifecyclemanager.LifecycleManager
}

func New(h string, p int, lm lifecyclemanager.LifecycleManager) CalculatorInvoker {
	ior := shared.IOR{Host: h, Port: p}
	inv := CalculatorInvoker{Ior: ior, lm: lm}
	return inv
}

func (i CalculatorInvoker) Invoke(b []byte) []byte {
	m := marshaller.Marshaller{}
	miopPacket := miop.Packet{}

	// Unmarshall miop packet
	miopPacket = m.Unmarshall(b)

	// Extract request from publisher
	r := miop.ExtractRequest(miopPacket)

	_p1 := float64(r.Params[0].(float64))
	_p2 := float64(r.Params[1].(float64))

	// Get instance from pool
	c := i.lm.GetObject()

	// Prepare reply
	var params []interface{}

	switch r.Op {
	case "Som":
		params = append(params, c.Som(_p1, _p2))
		i.lm.ReleaseObject(c)
	default:
		log.Fatal("Invoker:: Operation '" + r.Op + "' is unknown:: ")
	}

	// Create miop reply packet
	miop := miop.CreateReplyMIOP(params)

	// Marshall miop packet
	b = m.Marshall(miop)

	return b
}
