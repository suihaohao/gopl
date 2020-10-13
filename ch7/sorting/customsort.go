//Learning URL https://books.studygolang.com/gopl-zh/
package sorting

import "sort"

type customSort struct {
	t []*Track
	less func(x, y *Track)bool
}

func (x customSort)Len() int  {
	return len(x.t)
}
func (x customSort)Less(i, j int) bool  {
	return x.less(x.t[i], x.t[j])
}
func (x customSort)Swap(i, j int)  {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func useCustomSort()  {
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title{
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length{
			return x.Length < y.Length
		}
		return false
	}})
}