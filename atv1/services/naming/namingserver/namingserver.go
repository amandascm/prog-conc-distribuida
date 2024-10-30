package main

import (
	"fmt"
	naminginvoker "test/atv1/services/naming/invoker"
	"test/shared"
)

func main() {

	go namingServer()

	fmt.Println("'Servidor de Nomes' em execução...")
	fmt.Scanln()
}

func namingServer() {
	// Start messagingservice invoker
	i := naminginvoker.New(shared.LocalHost, shared.NamingPort)
	go i.Invoke()
}
