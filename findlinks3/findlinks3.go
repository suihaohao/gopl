//Learning URL https://books.studygolang.com/gopl-zh/
package findlinks3

import (
	"fmt"
	"gopl/links"
	"log"
)

func breadthFirst(f func(item string) []string, worklist []string)  {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items{
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil{
		log.Print(err)
	}
	return list
}

func RunBreadthFirst()  {
	breadthFirst(crawl,[]string{"https://golang.org"})
}
