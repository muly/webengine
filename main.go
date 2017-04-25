package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/welcome/", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}

type num struct {
	a int
	b int
}

func (n num) add() int {

	return n.a + n.b

}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!!!"))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	// localhost:8080/welcome/ram -> hello ram, welcome to team.

	u := r.URL.RequestURI()

	w.Write([]byte(u))
}
