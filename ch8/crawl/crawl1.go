//Learning URL https://books.studygolang.com/gopl-zh/
package crawl

import (
	"fmt"
	"gopl/links"
	"log"
	"os"
)

func crawl1(url string) []string  {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func RunCrawl1()  {
	log.Println(os.Args)
	worklist := make(chan []string)
	go func() {
		worklist <- os.Args[1:]
	}()
	seen := make(map[string]bool)
	for list := range worklist{
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl1(link)
				}(link)
			}
		}
	}
}

