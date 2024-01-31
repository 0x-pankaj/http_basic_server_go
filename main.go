package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("starting server \n")

	http.HandleFunc("/about", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "about page")
	})

	http.HandleFunc("/news", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "about news")
	})

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				w.WriteHeader(404)
				w.Write([]byte("404 - Not Found ! \n"))
				return
			}

			fmt.Fprintln(w, "home page!")

		})
	*/

	//handle files
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)

	//handle form
	http.HandleFunc("/form", handleForm)

	//handle query
	http.HandleFunc("/query", queryHandle)

	//handle image
	http.HandleFunc("/image", handleImage)

	log.Println("listening...")
	log.Fatal(http.ListenAndServe(":8090", nil))

}

func handleImage(w http.ResponseWriter, r *http.Request) {
	buf, err := os.ReadFile("./images/google-logo-1.png")

	if err != nil {
		// fmt.Println("error: ", err)
		log.Fatal(err)
	}
	w.Header().Set("content-type", "image/png")
	w.Write(buf)

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./public/form.html")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		// name := r.Form.Get("name")
		name := r.FormValue("name")
		address := r.Form.Get("address")
		fmt.Println(name, address)

		fmt.Fprintf(w, "%s address is %s", name, address)

	default:
		{
			fmt.Fprintf(w, "only GET and POST method is supported")
		}

	}
}

func queryHandle(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["name"]
	var name string
	if ok {
		name = keys[0]
	}
	fmt.Fprintf(w, "hello %s\n", name)

}
