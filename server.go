package main

import (
	"fmt"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	r.Proto = "HTTP/1.0"
	fmt.Println("form", r.Form)
	fmt.Println("scheme", r.URL.Scheme)
	//w.Header().Set("HTTP/1.0 200 OK\n\n")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
