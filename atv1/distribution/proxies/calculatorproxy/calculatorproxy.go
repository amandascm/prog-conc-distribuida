package calculatorproxy

import (
	"test/atv1/distribution/core"
	"test/atv1/distribution/requestor"
	"test/shared"
)

type CalculatorProxy struct {
	Ior   shared.IOR
	_Core core.Core
}

func New(i shared.IOR) CalculatorProxy {
	r := CalculatorProxy{Ior: i}
	return r
}

func (p *CalculatorProxy) Som(p1, p2 int) int {

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

func (h *CalculatorProxy) Dif(p1, p2 int) int {

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

func (h *CalculatorProxy) Mul(p1, p2 int) int {

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

func (h *CalculatorProxy) Div(p1, p2 int) int {

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
