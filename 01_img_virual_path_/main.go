package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("img"))))
	// http.Handle("/", http.FileServer(http.Dir("img")))
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":3333", nil)

}
func home(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "THis is home page")

}
func contact(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprintf(response, "<img src=\"/20201217_124208.jpg\" alt=\"my\" width=\"500\" height=\"600\">")

}
func about(response http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(response, "THis is about page")

}
