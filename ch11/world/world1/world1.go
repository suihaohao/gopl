//Learning URL https://books.studygolang.com/gopl-zh/
package world1

func IsPalindrome(s string) bool  {
	for i := range s{
		if s[i] != s [len(s) - 1-i]{
			return false
		}
	}
	return true

	//n := len(s)/2
	//for i := 0; i < n; i++ {
	//	if s[i] != s[len(s)-1-i] {
	//		return false
	//	}
	//}
	//return true
}
