package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()      
	
	// print in server
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println("------------")
	
	// print in client
	fmt.Fprintf(w, "Wow, you find this.")
	for k, v := range r.Form{
		fmt.Fprintf(w, k)
		fmt.Fprintf(w, " = ") 
		fmt.Fprintf(w, strings.Join(v, ""))
		fmt.Fprintf(w, ", ")
	}
}

func main() {
	fmt.Println("Start running...")
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
