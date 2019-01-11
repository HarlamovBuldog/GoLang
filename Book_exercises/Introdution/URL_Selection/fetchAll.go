package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	prefix := "http://"
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			var s []string
			s = append(s, prefix)
			s = append(s, url)
			url = strings.Join(s, "")
		}
		go fetch(url, ch)
	}

	// open output file
	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)
	for range os.Args[1:] {
		//fmt.Println(<-ch)
		if _, err := w.WriteString(<-ch + "\n"); err != nil {
			panic(err)
		}
		//Getting from channel ch
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //sending to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() //Exception information leakage
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v:", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
