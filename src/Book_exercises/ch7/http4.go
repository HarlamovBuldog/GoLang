package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	//http.HandleFunc(db.list) string means type conversion,
	//not call of the function
	mux.Handle("/list", http.HandlerFunc(db.list))
	// these 2 strings can be shorten to
	// mux.Handle("/price", db.price)
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/update", http.HandlerFunc(db.update))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		priceStr := req.URL.Query().Get("price")
		priceFloat, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			fmt.Fprintf(w, "wrong price: %s\n", priceStr)
			return
		}
		db[item] = dollars(float32(priceFloat))
		fmt.Fprintf(w, "%s item with price %s was successfully added!\n", item, db[item])
		return
	}
	fmt.Fprintf(w, "item with name \"%s\" already exist!\nIf u wanted to update item use /update\n", item)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	oldPrice, ok := db[item]
	if ok {
		priceStr := req.URL.Query().Get("price")
		priceFloat, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			fmt.Fprintf(w, "wrong price: %s\n", priceStr)
			return
		}
		db[item] = dollars(float32(priceFloat))
		fmt.Fprintf(w, "%s item's price %s was successfully updated to %s!\n", item, oldPrice, db[item])
		return
	}
	fmt.Fprintf(w, "item with name \"%s\" doesn't exist!\n", item)
}
