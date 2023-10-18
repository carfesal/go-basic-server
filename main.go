package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func form(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" { // if the path is not /hello
		http.Error(w, fmt.Sprintf("URL NOT FOUND: %s", r.URL.Path), http.StatusNotFound) // send a 404 not found error
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprint(w, "Post successful \n")

	//getting the values of the form
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)

}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" { // if the path is not /hello
		http.Error(w, fmt.Sprintf("URL NOT FOUND: %s", r.URL.Path), http.StatusNotFound) // send a 404 not found error
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	fmt.Fprintf(w, fmt.Sprintf("Hello World in Go: %s!!!", runtime.Version())) // print hello
}

func main() {
	server := http.FileServer(http.Dir("./static")) // create a server checking the static folder

	http.Handle("/", server) // handle the route / and send it to the server

	http.HandleFunc("/form", form) // handle the route /form and send it to the <form function>

	http.HandleFunc("/hello", hello) // handle the route /hello and send it to the <helloHandler> function

	fmt.Printf("Starting server at port 8001\n")

	if err := http.ListenAndServe(":8001", nil); err != nil { // listen and serve the server at port 8080
		log.Fatal(err)
	}
}
