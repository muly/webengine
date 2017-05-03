package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type vamsi struct {
}

func (vamsi *vamsi) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]

	log.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write([]byte(data))
	} else {
		w.Write([]byte("404 vamsi - " + http.StatusText(404)))
	}
}
func main() {

	log.Println("server started vamsi")
	http.Handle("/", new(vamsi))
	http.ListenAndServe(":8080", nil)

}
