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
	"time"
)

//like hash table
var dispatch map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var path string
	path = r.URL.Path[1:]

	//dispatch[path] => ponteiro for function, status(nil, true)
	if h, ok := dispatch[path]; ok{
		h(w, r)
		return
	}

	//if not in dispatch

	conn := w.Connection
	//status-line
	w.WriteHeader(http.StatusOK)

	//time, used to cache
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	//name server and content-type
	w.Header().Set("Server", "Servidor-Cecilia")
	w.Header().Set("Content-type", "text/html")
	
	fmt.Fprintf(w, "oi mundo")
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("Proto", "HTTP/1.0")
	w.WriteHeader(http.StatusOK)

	//time, used to cache
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format("2016-10-23"))

	//name server and content-type
	w.Header().Set("Server", "Servidor-Cecilia")
	w.Header().Set("Content-type", "text/html")
	
	fmt.Println("form", r.Form)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler:  &myHandler{},
	}

	dispatch = make(map[string]func(http.ResponseWriter, *http.Request))

	dispatch["cecilia"] = handler
	//dispatch["favicon.ico"] = handler

	server.ListenAndServe()
}
