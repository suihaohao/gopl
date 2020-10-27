//Learning URL https://books.studygolang.com/gopl-zh/
package bank

import (
	"sync"
)

var(
	mu sync.RWMutex
	balance3 int
)

func Deposit3(amount int)  {
	mu.RLock()
	defer mu.RUnlock()
	balance3 = balance3 + amount
}

func Balance3() int {
	mu.RLock()
	defer mu.RUnlock()
	b := balance3
	return b
}
