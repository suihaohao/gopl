//Learning URL https://books.studygolang.com/gopl-zh/
package sum

func sum(vals ...int) int {
	total := 0
	for _, val := range vals{
		total += val
	}
	return total
}
