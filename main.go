package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/welcome/", welcomeHandler)
        http.HandleFunc("/welcome/ram", welcomeHandler)
	http.HandleFunc("/welcome/vamsi", vamsiHandler)
	http.ListenAndServe(":8080", nil)
}




func helloHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!!!"))
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	// localhost:8080/welcome/ram -> hello ram, welcome to team.

	//u := r.URL.RequestURI()

	w.Write([]byte("hi ram, welcome to the team"))
}


func vamsiHandler(x http.ResponseWriter, y *http.Request){

        x.Write([]byte("hi vamsi, welcome to the team"))



}






