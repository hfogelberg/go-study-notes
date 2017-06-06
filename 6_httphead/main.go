/**
Prints entire head of a web site.
Could be used for debug logging if rewritten as middleware.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, %s, %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddre = %q", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}

		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", v, k)
		}
	}
}
