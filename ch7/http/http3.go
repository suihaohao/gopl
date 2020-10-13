//Learning URL https://books.studygolang.com/gopl-zh/
package http

import (
	"fmt"
	"net/http"
)

var Mux = http.NewServeMux()

func init(){
	db := Database{"shoes": 50, "socks": 5}
	Mux.HandleFunc("/list", db.list)
	Mux.HandleFunc("/price", db.price)
}

func (db Database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db Database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}