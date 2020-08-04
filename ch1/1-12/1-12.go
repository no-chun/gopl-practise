package main

import (
	"gopl-practise/ch1/lissajous"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}
	var cycle int
	for k, v := range request.Form {
		if k == "cycle" {
			var err error
			cycle, err = strconv.Atoi(strings.Join(v, ""))
			if err != nil {
				log.Print("cycle must be int")
			}
		}
	}
	lissajous.Lissajous(writer, cycle)
}
