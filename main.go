package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gurdeep/crud/routers"
)

func main() {
	//get a router
	router := routers.GetMuxRouter()

	//Listen and Serve
	//If we do not want to use gorilla mux, router can be nil
	/*
		http.Handle("/foo", fooHandler)
		http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
		http.ListenAndServe(":8000", nil)
	*/
	fmt.Println("CRUD project started....listening on port 8000!")
	log.Fatal(http.ListenAndServe(":8000", router))
}
