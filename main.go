package main

import (
	"fmt"
	"gopl/fibonacci"
)

func main() {
	//lissajous.RunLissajous()
	//fetch_url.RunFetch()
	//web_service.RunWebService()
	//for i:=0.0; i <5*2*math.Pi; i ++ {
	//	fmt.Println(math.Sin(i), i)
	//}
	for i := 1; i < 10; i++ {
		fmt.Println(fibonacci.GetNthFib(i))
	}
}
