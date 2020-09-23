package main

import (
	"fmt"
	"gopl/reverse"
)

func main() {
	//lissajous.RunLissajous()
	//fetch_url.RunFetch()
	//web_service.RunWebService()
	fmt.Println(reverse.ReverseArr(&[]interface{}{1,2,3,4, "a", "b","c"}))
}
