package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	servers := parse(os.Args[1:])
	for _, serv := range servers {
		conn, err := net.Dial("tcp", serv.addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go func(s *server) {
			sc := bufio.NewScanner(conn)
			for sc.Scan() {
				serv.time = sc.Text()
			}
		}(serv)
	}

	for {
		fmt.Printf("\n")
		for _, server := range servers {
			fmt.Printf("%s: %s\n", server.name, server.time)
		}
		fmt.Print("------------------")

		time.Sleep(time.Second)
	}
}

type server struct {
	name string
	addr string
	time string
}

func parse(args []string) []*server {
	var servers []*server
	for _, arg := range args {
		s := strings.SplitN(arg, "=", 2)
		servers = append(servers, &server{s[0], s[1], ""})
	}
	return servers
}
