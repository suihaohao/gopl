//Learning URL https://books.studygolang.com/gopl-zh/
package findlinks

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func RunOutline() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = n.NextSibling {
		outline(stack, c)
	}
}
