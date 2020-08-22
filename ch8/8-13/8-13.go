package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			var names []string
			for c := range clients {
				names = append(names, c.name)
			}
			cli.ch <- fmt.Sprintf("%d arrival: %v", len(names), names)
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	input := bufio.NewScanner(conn)
	timer := time.NewTimer(5 * time.Minute)

	go func() {
		<-timer.C
		_ = conn.Close()
		return
	}()

	for input.Scan() {
		timer.Reset(5 * time.Minute)
		messages <- who + ": " + input.Text()
	}

	leaving <- client{ch, who}
	messages <- who + " has left"
	_ = conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		_, _ = fmt.Fprintln(conn, msg)
	}
}
