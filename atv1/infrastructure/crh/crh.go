package crh

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
	"test/shared"
)

type CRH struct {
	Host       string
	Port       int
}

func NewCRH(h string, p int) *CRH {
	r := new(CRH)

	r.Host = h
	r.Port = p

	return r
}

func (crh *CRH) SendReceive(msgToServer []byte) []byte {
	var err error
	var conn net.Conn

	s := 0
	msgFromServer := []byte{}

	for {
		switch s {
		case 0: // open connection
			for i := 0; i < shared.MaxConnectionAttempts; i++ {
				conn, err = net.Dial("tcp", crh.Host+":"+strconv.Itoa(crh.Port))
				if err == nil {
					s = 1
					break
				} else {
					if i == shared.MaxConnectionAttempts-1 {
						log.Fatal("CRH 0:: Number Max of attempts achieved...")
					}
				}
			}
		case 1:
			// 2: send message's size
			sizeMsgToServer := make([]byte, 4)
			l := uint32(len(msgToServer))
			binary.LittleEndian.PutUint32(sizeMsgToServer, l)
			_, err = conn.Write(sizeMsgToServer)
			if err != nil {
				log.Fatalf("CRH 1:: %s", err)
			}
			s = 2
		case 2:
			// 3: send message
			_, err = conn.Write(msgToServer)
			if err != nil {
				log.Fatalf("CRH 2:: %s", err)
			}
			s = 3
		case 3:
			// 4: receive message's size
			sizeMsgFromServer := make([]byte, 4)
			_, err = conn.Read(sizeMsgFromServer)
			if err != nil {
				log.Fatalf("CRH 3:: %s", err)
			}
			sizeFromServerInt := binary.LittleEndian.Uint32(sizeMsgFromServer)

			//5: receive reply
			msgFromServer = make([]byte, sizeFromServerInt)
			_, err = conn.Read(msgFromServer)
			if err != nil {
				log.Fatalf("CRH 4:: %s", err)
			}
			s = 4
		case 4:
			conn.Close()
			return msgFromServer
		}
	}
}
