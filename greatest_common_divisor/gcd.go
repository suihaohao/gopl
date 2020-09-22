//Learning URL https://books.studygolang.com/gopl-zh/
package greatest_common_divisor

import "fmt"

func GetGcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
		fmt.Println(x, y)
	}
	return x
}
