package web_service

import (
	"fmt"
	lissajous2 "gopi/lissajous"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func RunWebService() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", lissajous)
	http.HandleFunc("/count", counter) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
func lissajous(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	var cycles float64 = 5
	for k, v := range r.Form {
		if k == "cycles" {
			fmt.Println(v)
			c, err := strconv.Atoi(v[0])
			if err != nil {
				log.Println(err)
			} else {
				cycles = float64(c)
			}
		}
	}
	lissajous2.LissajousForWeb(w, cycles)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)

	}
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
