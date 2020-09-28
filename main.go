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
	//fmt.Println(html.UnescapeString("https:&#47;&#47;www.footlocker.co.uk&#47;INTERSHOP&#47;web&#47;FLE&#47;Footlocker-Footlocker_GB-Site&#47;en_GB&#47;-&#47;GBP&#47;BarclaycardSmartpayCheckoutCallback-Return"))
}
