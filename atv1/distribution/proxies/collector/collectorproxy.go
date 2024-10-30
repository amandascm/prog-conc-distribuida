package collectorproxy

import (
	"test/atv1/distribution/core"
	"test/atv1/distribution/requestor"
	"test/shared"
)

type CollectorProxy struct {
	Ior   shared.IOR
	_Core core.Core
}

func New(i shared.IOR) CollectorProxy {
	r := CollectorProxy{Ior: i}
	return r
}

func (p *CollectorProxy) Log(message string) {

	// 1. Configure input parameters
	params := make([]interface{}, 1)
	params[0] = message

	// Configure remote request
	req := shared.Request{Op: "Log", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: p.Ior, Request: req}

	// 3. Invoke Requestor
	requestor := requestor.Requestor{}
	requestor.Invoke(inv)
}

func (p *CollectorProxy) Som(p1, p2 int) int {

	// 1. Configure input parameters
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Configure remote request
	req := shared.Request{Op: "Som", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: p.Ior, Request: req}

	// 3. Invoke Requestor
	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	//4. Return something to the publisher
	return int(r.Rep.Result[0].(float64))
}

func (h *CollectorProxy) Dif(p1, p2 int) int {

	// 1. Configure input parameters
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Configure remote request
	req := shared.Request{Op: "Dif", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// 3. Invoke Requestor
	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	//4. Return something to the publisher
	return int(r.Rep.Result[0].(float64)) // TODO
}

func (h *CollectorProxy) Mul(p1, p2 int) int {

	// 1. Configure input parameters
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Configure remote request
	req := shared.Request{Op: "Mul", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// 3. Invoke Requestor
	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	//4. Return something to the publisher
	return int(r.Rep.Result[0].(float64)) // TODO
}

func (h *CollectorProxy) Div(p1, p2 int) int {

	// 1. Configure input parameters
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2

	// Configure remote request
	req := shared.Request{Op: "Div", Params: params}

	// Prepare invocation to Requestor
	inv := shared.Invocation{Ior: h.Ior, Request: req}

	// 3. Invoke Requestor
	requestor := requestor.Requestor{}
	r := requestor.Invoke(inv)

	//4. Return something to the publisher
	return int(r.Rep.Result[0].(float64)) // TODO
}
