/**
Fetch several urls concurrently.
It will take as long as the longest fetch
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive response from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	s := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs, %7d %s", s, nbytes, url)
}
