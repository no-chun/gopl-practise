package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	defer conn.Close()

	input := bufio.NewScanner(conn)

	ch := make(chan struct{})
	go func() {
		for input.Scan() {
			ch <- struct{}{}
		}
	}()

	for {
		select {
		case <-time.After(10 * time.Second):
			return
		case <-ch:
			go echo(conn, input.Text(), 1*time.Second)
		}
	}
}

func echo(w io.Writer, text string, delay time.Duration) {
	_, _ = fmt.Fprintln(w, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	_, _ = fmt.Fprintln(w, "\t", strings.ToLower(text))
	time.Sleep(delay)
	_, _ = fmt.Fprintln(w, "\t", strings.ToLower(text))
}
