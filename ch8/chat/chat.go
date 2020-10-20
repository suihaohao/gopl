//Learning URL https://books.studygolang.com/gopl-zh/
package chat

import (
	"log"
	"net"
)

func RunChat()  {
	listener, err :=net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil{
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

