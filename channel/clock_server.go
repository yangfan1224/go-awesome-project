package channel

import (
	"net"
	"log"
	"io"
	"time"
)

func handlConn(c net.Conn){
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}

}

func StartServer(){
	listener, err := net.Listen("tcp","localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)// such as connection close
			continue
		}

		go handlConn(conn)
	}
}
