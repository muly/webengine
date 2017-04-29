package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("server started")

	http.HandleFunc("/", templateHandler)

	http.ListenAndServe(":8090", nil)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {

	//template.New("name")
	t1 := template.Must(template.ParseFiles(
		"public/table.html",
		//"public/welcome.html",
	))

	err := t1.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
