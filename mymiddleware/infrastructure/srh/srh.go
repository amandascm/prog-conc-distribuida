package srh

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type SRH struct {
	Host       string
	Port       int
	Connection net.Conn
}

var ln net.Listener

// var conn net.Conn
var err error

func NewSRH(h string, p int) *SRH {
	r := new(SRH)

	r.Host = h
	r.Port = p
	r.Connection = nil

	return r
}

func (srh *SRH) Receive() []byte {

	// 1: create listener & accept connection
	if srh.Connection == nil {
		ln, err = net.Listen("tcp", srh.Host+":"+strconv.Itoa(srh.Port))
		if err != nil {
			log.Fatalf("SRH:: %s", err)
		}

		srh.Connection, err = ln.Accept()
		if err != nil {
			log.Fatalf("SRH:: %s", err)
		}
	}

	// 2: receive message's size
	size := make([]byte, 4)
	_, err = srh.Connection.Read(size)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			return nil
		} else {
			log.Fatalf("SRH:: %s", err)
		}
	}
	sizeInt := binary.LittleEndian.Uint32(size)

	// 3: receive message
	msg := make([]byte, sizeInt)
	_, err = srh.Connection.Read(msg)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			return nil
		} else {
			log.Fatalf("SRH:: %s", err)
		}
	}

	return msg
}

func (srh *SRH) Send(msgToClient []byte) {

	// 1. Check availability of connection
	if srh.Connection == nil {
		fmt.Println("SRH:: Connection not opened")
		os.Exit(0)
	}

	// 2: send message's size
	size := make([]byte, 4)
	l := uint32(len(msgToClient))
	binary.LittleEndian.PutUint32(size, l)
	_, err = srh.Connection.Write(size)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			return
		} else {
			log.Fatalf("SRH:: %s", err)
		}
	}

	// 3: send message
	_, err = srh.Connection.Write(msgToClient)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			srh.Connection.Close()
			ln.Close()
			return
		} else {
			log.Fatalf("SRH:: %s", err)
		}
	}

	// 4: close connection
	srh.Connection.Close()
	srh.Connection = nil

	ln.Close()
}
