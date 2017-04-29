package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type MyHandler struct {
}

func (person *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Path[1:]
	log.Println(a)

	data, err := ioutil.ReadFile(string(a))
	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 vamsi- " + http.StatusText(404)))
	}
}
func main() {

	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080", nil)

}

//func hellohandler(rw http.ResponseWriter, req *http.Request) {
//rw.Write([]byte("hello vamsi, sucessfully launched the server"))
//}
//func vamsihandler(w http.ResponseWriter, r *http.Request) {
//fmt.Fprintf(w, "Hii %s \n how may i help you today", r.URL.Path[3:])
//}
