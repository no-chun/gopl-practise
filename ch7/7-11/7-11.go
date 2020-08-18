package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var db = database{"shoes": 50, "socks": 5}

func main() {
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/creat", db.creat)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

var mutex sync.Mutex

type dollars float64
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	_, _ = fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item, price := r.URL.Query().Get("item"), r.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "params err: %s\n", err)
	}
	mutex.Lock()
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "%s isn't exist\n", item)
	} else {
		db[item] = dollars(p)
	}
	mutex.Unlock()
}

func (db database) creat(w http.ResponseWriter, r *http.Request) {
	item, price := r.URL.Query().Get("item"), r.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "params err: %s\n", err)
	}
	mutex.Lock()
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "%s isn exist\n", item)
	} else {
		db[item] = dollars(p)
	}
	mutex.Unlock()
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "%s isn't exist\n", item)
	} else {
		mutex.Lock()
		delete(db, item)
		mutex.Unlock()
	}
}
