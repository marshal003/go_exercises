package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	urls := map[string]string{
		"/docs":      "https://godoc.org",
		"/yaml-godo": "https://godoc.org/gopkg.in/yaml.v2",
	}

	handlerFunc := mapHandler(urls, mux)
	http.ListenAndServe(":8080", handlerFunc)
}

func mapHandler(urls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := urls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
