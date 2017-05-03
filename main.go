package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("server started")

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/welcome/", welcomeHandler)
	http.HandleFunc("/welcome/ram", ramhandler)
	http.HandleFunc("/welcome/vamsi", vamsihandler)
	http.ListenAndServe(":8080", nil)

	http.HandleFunc("/table", tableHandler)
	http.HandleFunc("/welcome", welcomeHandler)

	http.ListenAndServe(":8090", nil)
}

func tableHandler(w http.ResponseWriter, r *http.Request) {

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

	w.Write([]byte("welcome to v groups!!!!"))
}
func ramhandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hii %s, welcome to the team !", r.URL.Path[2:])
	w.Write([]byte("Hii ram, welcome to v groups!!!!"))
}
func vamsihandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hii vamsi, welcome to v groups !!!"))

	//template.New("name")
	t1 := template.Must(template.ParseFiles(
		"public/welcome.html",
	))

	err := t1.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
