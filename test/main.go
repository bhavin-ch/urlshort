package main

import (
	"fmt"
	"net/http"
	"github.com/bhavin-ch/urlshort"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	yaml := `
- path: /urlshort
  url: https://github.com/bhavin-ch/urlshort
- path: /urlshort-final
  url: https://github.com/bhavin-ch/urlshort/tree/master
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8888")
	err1 := http.ListenAndServe(":8888", yamlHandler)
	if err1 != nil {
		panic(err1)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	return mux
}
