//Learning URL https://books.studygolang.com/gopl-zh/
package du

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func RunDu3()  {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir2(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <- chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop :
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok{
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

func walkDir3(dir string, n *sync.WaitGroup, fileSizes chan<- int64)  {
	defer n.Done()
	for _, entry := range dirents3(dir){
		if entry.IsDir(){
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir3(subDir, n, fileSizes)
		}else {
			fileSizes <- entry.Size()
		}
	}
}


var sema = make(chan struct{}, 20)
func dirents3(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}