package channel

import (
	"net"
	"io"
	"log"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalf("mustCopy failed: %s", err)
	}
}
func SyncChannel(){
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func(){
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	if f, ok := conn.(*net.TCPConn); ok {
		f.CloseWrite()
	}
	//conn.Close()
	<- done
}
