//Learning URL https://books.studygolang.com/gopl-zh/
package crawl

import (
	"fmt"
	"gopl/links"
	"log"
	"os"
	"strings"
)

var tokens =make(chan struct{},20)
func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<- tokens
	if err != nil {
		log.Println(err)
	}
	return list
}

func RunCrawl2()  {
	worklist := make(chan []string)
	var n int
	n ++
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for ;n >0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link]{
				seen[link] = true
				n ++
				go func(link string) {
					worklist <- crawl1(link)
				}(link)
			}
		}
	}
}

func RunCrawl22() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	go func() {worklist <- os.Args[1:]}()
	for i:=0; i < 20; i ++ {
		go func() {
			for link := range unseenLinks{
				//if strings.HasPrefix(link, "https://gopl.io/"){
					foundLinks := crawl2(link)
					go func() {
						worklist <- foundLinks
					}()
				//}
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist{
		for _, link := range list{
			if strings.HasPrefix(link, "https://gopl.io/") {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		}
	}
}