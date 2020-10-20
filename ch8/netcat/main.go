//Learning URL https://books.studygolang.com/gopl-zh/
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main()  {
	netcat3()
}

func netcat1(){
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}
func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func netcat2()  {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func netcat3()  {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<- done
}