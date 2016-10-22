/*
|**********************************************************************;
* Projeto           : Simple HTTP web-server in Golang.
*
* Program name      : server.go
*
* Author            : Cecília Carneiro e Silva
*
* Date created      : 22/10/2016
*
* Purpose           : Server HTTP/1.0 using Golang, net/http.
*
* Uso               : $ go run server.go
*
|**********************************************************************;
*/

package main

import (
	"fmt"
        "net/http"
)

//like hash table
var dispatch map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var path string
	path = r.URL.Path[1:]
	if h, ok := dispatch[path]; ok{
		h(w, r)
		return
	}

	fmt.Fprintf(w, "oi mundo")
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	r.Proto = "HTTP/1.0"
	fmt.Println("form", r.Form)
	fmt.Println("scheme", r.URL.Scheme)
	//w.Header().Set("HTTP/1.0 200 OK\n\n")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler:  &myHandler{},
	}

	dispatch = make(map[string]func(http.ResponseWriter, *http.Request))

	dispatch["/cecilia"] = handler

	server.ListenAndServe()
}
