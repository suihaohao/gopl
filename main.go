package main

import (
	"fmt"
	"gopl/shuffle"
	"time"
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
	//findlinks3.RunBreadthFirst()
	//log.Fatal(http.ListenAndServe("localhost:8000", myHttp.Mux))
	//reverb.RunReverbHandleConn()
	a := []int{1,2,3,4}
	for _,m := range a {
		go func() {
			fmt.Println(m)
		}()
	}
	time.Sleep(time.Second)
}

