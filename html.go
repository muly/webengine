package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type vamsi struct {
}

func (vamsi *vamsi) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]

	fmt.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write([]byte(data))
	} else {

		w.Write([]byte("404 vamsi - " + http.StatusText(404)))
	}
}
func main() {

	fmt.Println("server started")
	http.Handle("/", new(vamsi))
	http.ListenAndServe(":8080", nil)

}
