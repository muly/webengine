package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started")
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello!!!"))
}
