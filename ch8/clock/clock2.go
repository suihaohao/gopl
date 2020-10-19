//Learning URL https://books.studygolang.com/gopl-zh/
package clock

import (
	"log"
	"net"
)

func RunHandelConn2()  {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}