package web_service

import (
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func RunWebService() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", lissajous)
	http.HandleFunc("/surface", surface)
	http.HandleFunc("/count", counter) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
