//Learning URL https://books.studygolang.com/gopl-zh/
package tempconv

import (
	"flag"
	"fmt"
)
var temp = CelsiusFlag("temp", 20.0, "the temperature")

func RunTempFlag()  {
	flag.Parse()
	fmt.Println(*temp)
}
