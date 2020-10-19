//Learning URL https://books.studygolang.com/gopl-zh/
package pipeline

import (
	"fmt"
)

func pipeline1()  {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x:= 0;;x++{
			naturals <- x
		}
	}()
	go func() {
		for  {
			x, ok:= <- naturals
			if !ok {
				break
			}
			squares <- x*x
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
	}
}
