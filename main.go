package main

import (
	"fmt"
	"log"
	"net/http"
)

func loggingDecorator(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL)
		f(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	rawurl := ""
	if url, ok := r.URL.Query()["url"]; ok {
		rawurl = url[0]
	}
	sid, err := AddPage(rawurl)
	if err != nil {
		log.Printf("[handler] error = %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s -> %s\n", sid, rawurl)
}

func main() {
	http.HandleFunc("/shortn", loggingDecorator(handler))
	log.Printf("Server starts listening :8080")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
