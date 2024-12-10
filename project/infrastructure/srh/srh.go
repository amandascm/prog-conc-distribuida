package srh

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"

	"test/project/distribution/invokers"
)

type SRH struct {
	Host    string
	Port    int
	Ln      net.Listener
	invoker invokers.Invoker
}

// var conn net.Conn
var err error

func NewSRH(h string, p int) *SRH {
	r := new(SRH)

	r.Host = h
	r.Port = p
	// 1: create listener & accept connection
	r.Ln, err = net.Listen("tcp", h+":"+strconv.Itoa(p))
	if err != nil {
		log.Fatalf("SRH 0:: %s", err)
	}

	return r
}

func NewWithInvoker(h string, p int, invoker invokers.Invoker) *SRH {
	r := new(SRH)

	r.Host = h
	r.Port = p
	// 1: create listener & accept connection
	r.Ln, err = net.Listen("tcp", h+":"+strconv.Itoa(p))
	if err != nil {
		log.Fatalf("SRH 0:: %s", err)
	}

	r.invoker = invoker

	return r
}

func (srh *SRH) Serve() {
	for {
		msg, conn := srh.Receive()
		go func(conn net.Conn, msg []byte) {
			response := srh.invoker.Invoke(msg)
			srh.Send(conn, response)
		}(conn, msg)
	}
}

func (srh *SRH) Receive() ([]byte, net.Conn) {
	connection, err := srh.Ln.Accept()
	if err != nil {
		log.Fatalf("SRH 1:: %s", err)
	}

	// 2: receive message's size
	size := make([]byte, 4)
	_, err = connection.Read(size)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			connection.Close()
			return nil, nil
		} else {
			log.Fatalf("SRH 2:: %s", err)
		}
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	// 3: receive message
	msg := make([]byte, sizeInt)
	_, err = connection.Read(msg)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			connection.Close()
			return nil, nil
		} else {
			log.Fatalf("SRH 3:: %s", err)
		}
	}
	return msg, connection
}

func (srh *SRH) Send(conn net.Conn, msgToClient []byte) {

	// 2: send message's size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	_, err = conn.Write(size)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			conn.Close()
			return
		} else {
			log.Fatalf("SRH 4:: %s", err)
		}
	}

	// 3: send message
	_, err = conn.Write(msgToClient)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			conn.Close()
			// srh.Ln.Close()
			return
		} else {
			log.Fatalf("SRH 5:: %s", err)
		}
	}
}
