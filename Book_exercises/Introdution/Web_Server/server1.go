//minimal "echo"-server

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/", handler)	//each request calls handler
	log.Fatal(http.ListenAndServe(":8000", nil))
}

//Handler returns component of path from url-request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}