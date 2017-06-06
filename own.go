package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type nani struct {
}

func (handle *nani) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write([]byte(data))
	} else {
		w.Write([]byte("404 Page not found" + http.StatusText(404)))
	}

}

func main() {
	log.Println("Server started")
	http.Handle("/", new(nani))
	http.ListenAndServe(":8080", nil)
}
