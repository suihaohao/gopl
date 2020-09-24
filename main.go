package main

import (
	"gopl/shuffle"
)

func main() {
	//lissajous.RunLissajous()
	//fetch_url.RunFetch()
	//web_service.RunWebService()
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i + 1
	}
	arr = shuffle.ShuffleArrInt(arr)
}
