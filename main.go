package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("server started")

	http.HandleFunc("/table", tableHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	http.ListenAndServe(":8090", nil)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {

	//template.New("name")
	t1 := template.Must(template.ParseFiles(
		"public/table.html",
	))

	err := t1.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}


func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	//template.New("name")
	t1 := template.Must(template.ParseFiles(
		"public/welcome.html",
	))

	err := t1.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
