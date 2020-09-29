//Learning URL https://books.studygolang.com/gopl-zh/
package findlinks2

import (
	"fmt"
	"golang.org/x/net/html"
	"gopl/findlinks"
	"net/http"
)

func findLinks(url string)([]string, error)  {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK{
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil{
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return findlinks.Visit(nil, doc), nil
}

func CountWordAndImages(url string)(words, images int, err error)  {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {
	if n.NextSibling != nil {
		w, i := countWordsAndImages(n.NextSibling)
		words += w
		images += i
		return
	}
	if n.Type == html.ElementNode && n.Data == "image" {
		words+=1
		//images+=1
	}
	return
}