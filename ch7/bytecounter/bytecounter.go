//Learning URL https://books.studygolang.com/gopl-zh/
package bytecounter
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error)  {
	*c += ByteCounter(len(p))
	return len(p), nil
}

