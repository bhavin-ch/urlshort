package main

import (
	"fmt"
	"net/http"
	"github.com/bhavin-ch/urlshort"
)

func main() {
	mux := defaultMux()
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	return mux
}
