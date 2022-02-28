package main

import (
	"fmt"
	"net/http"
	"handler"
	"io/ioutil"
	"log"
)

func main() {

	data, err := ioutil.ReadFile("routes.yaml")

	if err != nil {
		log.Fatal(err)
	}

	fallback_handler := func(w http.ResponseWriter, r *http.Request) {

		http.Redirect(w, r, "http://www.google.com", http.StatusFound)

	}

	//http://localhost:8080/  You can test the server in this address.      

	redirects, err := handler.YAMLParser(data)

	if err != nil {

		log.Fatal(err)

	}

	html_request := handler.Redirect(redirects, fallback_handler)

	fmt.Printf("Server started at 8080 \n")

	if err := http.ListenAndServe(":8080", html_request); err != nil {
	
		log.Fatal(err)
	}

}
