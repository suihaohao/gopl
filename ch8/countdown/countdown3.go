//Learning URL https://books.studygolang.com/gopl-zh/
package countdown

import (
	"fmt"
	"os"
	"time"
)

func countdown3()  {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(time.Second * 1)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// todo
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}
