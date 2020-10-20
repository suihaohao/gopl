//Learning URL https://books.studygolang.com/gopl-zh/
package countdown

import (
	"fmt"
	"time"
)

func countdown1()  {
	tick := time.Tick(1 * time.Second)
	for countdown := 10;countdown>0;countdown--{
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
