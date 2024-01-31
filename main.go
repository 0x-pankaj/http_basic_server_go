package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("starting server \n")

	http.HandleFunc("/about", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "about page")
	})

	http.HandleFunc("/news", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "about news")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(404)
			w.Write([]byte("404 - Not Found ! \n"))
			return
		}

		fmt.Fprintln(w, "home page!")

	})

	log.Println("listening...")
	log.Fatal(http.ListenAndServe(":8090", nil))

}
