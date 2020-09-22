//Learning URL https://books.studygolang.com/gopl-zh/
package fibonacci

func GetNthFib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}