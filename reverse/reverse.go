//Learning URL https://books.studygolang.com/gopl-zh/
package reverse

import "strings"

func ReverseArr(s *[]interface{}) []interface{} {
	for i := 0; i < len(*s)/2; i++ {
		(*s)[i], (*s)[len(*s)-i-1] = (*s)[len(*s)-i-1], (*s)[i]
	}
	//for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
	//	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	//}
	return *s
}

func ReverseString(s string) string {
	sArr := strings.Split(s, "")
	for i, j := 0, len(sArr)-1; i < j; i, j = i+1, j-1 {
		sArr[i], sArr[j] = sArr[j], sArr[i]
	}
	return strings.Join(sArr, "")
}
