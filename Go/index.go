package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Name  string
	Price int
}

func main() {

	var templates = template.Must(template.ParseFiles("index.html"))

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myProduct := Product{"นมสด", 500}
		templates.ExecuteTemplate(w, "index.html", myProduct)
		//http.ServeFile(w, r, "index.html")
	})

	userDB := map[string]int{
		"heng":   30,
		"golang": 10,
		"java":   40,
		"php":    50,
	}

	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "signup.html")
	})

	router.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})

	router.HandleFunc("/File", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "file.txt")
	})

	router.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := userDB[name]
		fmt.Fprintf(w, "My name is %s %d year old", name, age)
	}).Methods("GET")

	router.HandleFunc("/login", login)

	log.Fatal(http.ListenAndServe(":80", router))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Method:", r.Method)
	r.ParseForm()
	fmt.Println("Username:", r.Form["username"])
	fmt.Println("Password:", r.Form["password"])
}
