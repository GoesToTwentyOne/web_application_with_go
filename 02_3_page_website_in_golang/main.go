package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/request", request)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}
func init() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:Alex@tcp(127.0.0.1:3306)/hosting_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	fmt.Println("connection Ok")
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
	//method 1 name
	// name := req.FormValue("name")
	// company := req.FormValue("company")
	// email := req.FormValue("email")
	// fmt.Println(name, company, email)
	// fmt.Fprintf(response, "Received %s %s %s", name, company, email)
	//method 2
	req.ParseForm()
	for key, value := range req.Form {
		fmt.Println(key, value)
	}
	fmt.Fprintf(response, "OK")

}
