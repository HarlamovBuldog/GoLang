//Fetch.go displays a response to a
//request for an input URL

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	prefix := "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			var s []string
			s = append(s, prefix)
			s = append(s, url)
			url = strings.Join(s, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		bytes, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", os.Stdout, bytes)
		fmt.Println("\n" + resp.Status)
	}
}
