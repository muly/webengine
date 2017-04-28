package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("server started")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/welcome/", Handler)
    http.HandleFunc("/welcome/ram", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}




func helloHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!!!"))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	// localhost:8080/welcome/ram -> hello ram, welcome to team.

	u := r.URL.RequestURI()
	a := strings.Split(u, "/")

	//w.Write([]byte(a[2]), ("welcome to the team"))
	fmt.Fprintf(w, "hi %s, welcome to the team", a[2] )
}


func Handler(x http.ResponseWriter, y *http.Request){

        x.Write([]byte("hi Gopher, welcome to the team"))


}






