package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	cwd := "."
CLOSE:
	for scanner.Scan() {
		args := strings.Fields(scanner.Text())
		cmd := args[0]
		switch cmd {
		case "cd":
			if len(args) < 2 {
				_, _ = fmt.Fprintf(conn, "not enough argument")
			} else {
				cwd += "/" + args[1]
			}
		case "ls":
			if len(args) < 2 {
				_ = ls(conn, cwd)
			} else {
				path := args[1]
				if err := ls(conn, path); err != nil {
					_, _ = fmt.Fprint(conn, err)
				}
			}
		case "get":
			if len(args) < 2 {
				_, _ = fmt.Fprintf(conn, "not enough argument")
			} else {
				filename := args[1]
				content, err := ioutil.ReadFile(filename)
				if err != nil {
					_, _ = fmt.Fprint(conn, err)
				}
				_, _ = fmt.Fprintf(conn, "%s\n", content)
			}
		case "send":
			if len(args) < 2 {
				_, _ = fmt.Fprintf(conn, "not enough argument")
			} else {
				filename := args[1]
				file, err := os.Create(filename)
				if err != nil {
					_, _ = fmt.Fprint(conn, err)
				}

				cnt, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Fprint(conn, err)
				}

				var texts string
				for i := 0; i < cnt && scanner.Scan(); i++ {
					texts += scanner.Text() + "\n"
				}
				texts = strings.TrimSuffix(texts, "\n")
				_, _ = fmt.Fprint(file, texts)
				file.Close()
			}
		case "close":
			break CLOSE
		}
	}
}

func ls(w io.Writer, cwd string) error {
	files, err := ioutil.ReadDir(cwd)
	if err != nil {
		return err
	}
	for _, file := range files {
		_, _ = fmt.Fprintf(w, "%s\n", file.Name())
	}
	return nil
}
