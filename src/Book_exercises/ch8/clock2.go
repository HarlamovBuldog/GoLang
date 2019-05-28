// Clock2 presents TCP-server,
// which periodically shows time.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // for example, connection lost
			continue
		}
		//handleConn(conn)    // the only one connection processing
		go handleConn(conn) // for parallel processing
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05"))
		if err != nil {
			return // for example, client disconnection
		}
		time.Sleep(1 * time.Second)
	}
}
