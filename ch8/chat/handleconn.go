//Learning URL https://books.studygolang.com/gopl-zh/
package chat

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	message <- who + "has arrived"
	entering <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ": " + input.Text()
	}
	leaving <- ch
	message <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, "%s\n",msg)
	}
}
