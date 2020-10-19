//Learning URL https://books.studygolang.com/gopl-zh/
package pipeline

import (
	"fmt"
)

func pipeline3()  {

}

func counter(out chan<- int)  {
	for x :=0;x < 100;x++ {
		out <-x
	}
	close(out)
}

func square(out chan <- int, in <-chan int)  {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int)  {
	for v := range in{
		fmt.Println(v)
	}
}

func RunPipeline3()  {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go square(squares, naturals)
	printer(squares)
}