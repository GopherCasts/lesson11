package main

import (
	"flag"
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	host := flag.String("http", ":8080", "http host:port to listen on")
	flag.Parse()

	n := negroni.Classic()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Gophercasts!"))
	})

	n.UseHandler(mux)

	n.Run(*host)
}
