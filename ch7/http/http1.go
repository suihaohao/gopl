//Learning URL https://books.studygolang.com/gopl-zh/
package http

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String() string {return fmt.Sprintf("$%.2f", d)}

type Database map[string]dollars

func (db Database) ServeHTTP1(w http.ResponseWriter, req *http.Request)  {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
