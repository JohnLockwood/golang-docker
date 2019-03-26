package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body><p>Hello world.</p>"))
		if host, err := os.Hostname(); err == nil {
			w.Write([]byte("<p>Served from " + host + ".</p>"))
		}
		w.Write([]byte("</body></html>"))
	})

	http.ListenAndServe(":9000", r)
}
