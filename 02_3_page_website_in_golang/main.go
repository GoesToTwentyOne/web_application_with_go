package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/request", request)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}
func home(response http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())

	}
	ptmp.Execute(response, nil)

	// fmt.Fprintf(response, "Hello my home page")

}

func features(response http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(response, nil)

	// fmt.Fprintf(response, "Hello my home page")

}
func docs(response http.ResponseWriter, req *http.Request) {
	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(response, nil)

	// fmt.Fprintf(response, "Hello my home page")

}
func request(response http.ResponseWriter, req *http.Request) {
	//method name
	name := req.FormValue("name")
	company := req.FormValue("company")
	email := req.FormValue("email")
	fmt.Println(name, company, email)
	fmt.Fprintf(response, "Received %s %s %s", name, company, email)

}
