//Learning URL https://books.studygolang.com/gopl-zh/
package shuffle

import (
	"math/rand"
	"time"
)

func ShuffleArrInt(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	rand.Seed(time.Now().UnixNano())
	for i := range arr {
		idx := rand.Intn(len(arr))
		arr[i], arr[idx] = arr[idx], arr[i]
	}
	return arr
}
