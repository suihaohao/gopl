package fetch_url

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func RunFetch() {
	start := time.Now()
	//"https://golang.org",
	urls := [6]string{"http://gopl.io", "https://godoc.org", "http://gopl.io", "https://godoc.org", "http://gopl.io", "https://godoc.org"}
	ch := make(chan string)
	for _, url := range &urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
