//Learning URL https://books.studygolang.com/gopl-zh/
package reverb

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func RunReverbHandleConn()  {
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
		handleConn1(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\n", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn1(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}
