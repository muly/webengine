package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("server started")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/welcome/", vamsihandler)
	http.ListenAndServe(":8080", nil)

}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!!!"))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("welcome to v groups!!!!"))
}

func vamsihandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hii %s, welcome to the team !", r.URL.Path[9:])
}
