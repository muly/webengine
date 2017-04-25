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

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!!!"))
}

func welcomeHandler(rw http.ResponseWriter, req *http.Request) {

	//todo
	rw.Write([]byte("welcome!!!"))
}
