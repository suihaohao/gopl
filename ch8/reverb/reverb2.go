//Learning URL https://books.studygolang.com/gopl-zh/
package reverb

import (
	"bufio"
	"net"
	"time"
)

func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

