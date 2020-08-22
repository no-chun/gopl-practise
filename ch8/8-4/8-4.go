package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	var wg sync.WaitGroup

	input := bufio.NewScanner(conn)
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(conn, input.Text(), 1*time.Second)
		}()
	}

	wg.Wait()
	conn.Close()
}

func echo(w io.Writer, text string, delay time.Duration) {
	_, _ = fmt.Fprintln(w, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	_, _ = fmt.Fprintln(w, "\t", strings.ToLower(text))
	time.Sleep(delay)
	_, _ = fmt.Fprintln(w, "\t", strings.ToLower(text))
}
