//Learning URL https://books.studygolang.com/gopl-zh/
package reverse

import "strings"

func ReverseArr(s []interface{}) []interface{}{
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseString(s string) string{
	sArr :=strings.Split(s, "")
	for i, j := 0, len(sArr)-1; i < j; i, j = i+1, j-1 {
		sArr[i], sArr[j] = sArr[j], sArr[i]
	}
	return strings.Join(sArr, "")
}
