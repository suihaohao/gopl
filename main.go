package main

import (
	"fmt"
	"gopl/ch7/eval"
	"gopl/shuffle"
	"math"
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
	ru(func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

func ru(fn func(x, y float64) float64)  {
	fmt.Printf(fn)
}
