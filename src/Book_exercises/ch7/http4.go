// Exercise 7.11 and 7.12 realization from book page 236
// Create, update, delete functions were added to work with map
// "List" function was improved to show data in table using html/template
// Mutex was added to prevent simultaneous data changes
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

var goodsHTMLTable = template.Must(template.New("goodsHTMLTable").Parse(`
<h1>{{ len . }} items</h1>
<table border="1">
<tr style='text-align: left'>
	<th>Item</th>
	<th>Price</th>
</tr>
{{ range $item, $price := . }}
<tr>
	<td>{{ $item }}</td>
	<td>{{ $price }}</td>
</tr>
{{ end }}
</table>
`))

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
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) totalCount() int {
	return len(db)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	if err := goodsHTMLTable.Execute(w, &db); err != nil {
		log.Fatal(err)
	}
	/*
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	*/
	mu.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	price, ok := db[item]
	mu.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	_, ok := db[item]
	mu.Unlock()
	if !ok {
		priceStr := req.URL.Query().Get("price")
		priceFloat, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			fmt.Fprintf(w, "wrong price: %s\n", priceStr)
			return
		}
		mu.Lock()
		db[item] = dollars(float32(priceFloat))
		fmt.Fprintf(w, "%s item with price %s was successfully added!\n", item, db[item])
		mu.Unlock()
		return
	}
	fmt.Fprintf(w, "item with name \"%s\" already exist!\nIf u wanted to update item use /update\n", item)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	oldPrice, ok := db[item]
	mu.Unlock()
	if ok {
		priceStr := req.URL.Query().Get("price")
		priceFloat, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			fmt.Fprintf(w, "wrong price: %s\n", priceStr)
			return
		}
		mu.Lock()
		db[item] = dollars(float32(priceFloat))
		fmt.Fprintf(w, "%s item's price %s was successfully updated to %s!\n", item, oldPrice, db[item])
		mu.Unlock()
		return
	}
	fmt.Fprintf(w, "item with name \"%s\" doesn't exist!\n", item)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	price, ok := db[item]
	mu.Unlock()
	if ok {
		mu.Lock()
		delete(db, item)
		mu.Unlock()
		fmt.Fprintf(w, "Item %s with price %s was successfully deleted!\n", item, price)
		return
	}
	fmt.Fprintf(w, "Item with name \"%s\" doesn't exist!\n", item)
}
