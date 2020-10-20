//Learning URL https://books.studygolang.com/gopl-zh/
package du

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan int64)
func cancelled() bool {
	select {
	case <- done:
		return true
	default:
		return false
	}
}

func RunDu4()  {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots{
		n.Add(1)
		go walkDir4(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for  {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles ++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func walkDir4(dir string, n *sync.WaitGroup, fileSizes chan<- int64)  {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents4(dir){
		if entry.IsDir(){
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir4(subDir, n, fileSizes)
		}else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents4(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()
	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}
	return entries
}
