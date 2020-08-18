package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var db = database{"shoes": 50, "socks": 5}

func main() {
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type dollars float64
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./ch7/7-12/list.html"))
	if err := template.Execute(w, db); err != nil {
		log.Println(err)
	}
}
