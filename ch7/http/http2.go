//Learning URL https://books.studygolang.com/gopl-zh/
package http

import (
	"fmt"
	"net/http"
)

func (db Database) ServeHTTP2(w http.ResponseWriter, req *http.Request)  {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound)
	}
}
