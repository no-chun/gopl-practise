package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	w, c := CountingWriter(ioutil.Discard)
	fmt.Fprintf(w, "Hello!")
	fmt.Println(*c)
}

type Counter struct {
	w      io.Writer
	writen int64
}

func (c *Counter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.writen += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := Counter{w, 0}
	return &c, &c.writen
}
