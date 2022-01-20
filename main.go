package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}
type WorldHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rURL := r.RequestURI
	name := r.Form.Get("name")
	qValues := r.URL.Query()
	encQuers := qValues.Encode()
	qName := r.URL.Query().Get("name")
	msg := "Hello Handler! URL is " + rURL + " Name is " + name + " Query name is " + qName + " Encoded queries are " + encQuers
	w.Write([]byte(msg))
	//w.Write([]byte(name))
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
	w.Write([]byte("World handler"))
}

func main() {
	handler := MyHandler{}
	world := WorldHandler{}
	server := http.Server{Addr: "127.0.0.1:8087", Handler: &handler}
	http.Handle("/hello", &handler)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
