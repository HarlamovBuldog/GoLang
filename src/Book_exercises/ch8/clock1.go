// Clock1 presents TCP-server,
// which periodically shows time.
// Exercise 8.1 realization from book (port flag)
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var portFlag = flag.Int("port", 8000, "number of port to work with")

func main() {
	flag.Parse()
	strPortFlag := strconv.Itoa(*portFlag)
	listener, err := net.Listen("tcp", "localhost:"+strPortFlag)
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
