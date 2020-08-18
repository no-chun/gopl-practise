package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	reader := LimitReader(strings.NewReader("123456"), 1)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}

type LimitRead struct {
	reader io.Reader
	bytes  int64
}

func (r *LimitRead) Read(p []byte) (int, error) {
	if r.bytes <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > r.bytes {
		p = p[:r.bytes]
	}
	n, err := r.reader.Read(p)
	r.bytes -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitRead{
		reader: r,
		bytes:  n,
	}
}
