//Learning URL https://books.studygolang.com/gopl-zh/
package du

import (
	"flag"
	"time"
)

var verbose = flag.Bool("v", true, "show verbose progress messages")

func RunDu2()  {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
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
