package invokers

type Invoker interface {
	Invoke(b []byte) []byte
}
