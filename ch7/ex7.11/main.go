package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := newDataBase(map[string]dollars{"shoes": 50, "socks": 5})
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/delete", http.HandlerFunc(db.deleteItem))
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

type dollars float32

type database struct {
	v   map[string]dollars
	mux sync.Mutex
}

func newDataBase(v map[string]dollars) database {
	return database{
		v:   v,
		mux: sync.Mutex{},
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.v {
		fmt.Fprintf(w, "%s: %0.3f\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db.v[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%0.3f\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	db.mux.Lock()
	defer db.mux.Unlock()

	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")

	price, err := strconv.ParseFloat(priceStr, 32)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price %q is not a number\n", priceStr)
		return
	}

	db.v[item] = dollars(price)

	db.list(w, req)
}

func (db database) deleteItem(w http.ResponseWriter, req *http.Request) {
	db.mux.Lock()
	defer db.mux.Unlock()

	item := req.URL.Query().Get("item")

	delete(db.v, item)

	db.list(w, req)
}
