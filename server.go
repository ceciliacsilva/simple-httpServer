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
	"html/template"
	"strconv"
)

//show files
//http.ServeFile(w, r, "static/index.html") 

//like hash table
var dispatch map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}
type Add_result struct {
	Soma   float64
	Error  string
}


func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var path string
	path = r.URL.Path[1:]

	//dispatch[path] => ponteiro for function, status(nil, true)
	if h, ok := dispatch[path]; ok{
		h(w, r)
		return
	}

	//if not in dispatch

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

func adder(w http.ResponseWriter, r *http.Request){
	var soma  float64 = 0.0
	var error string  = " " 
	r.ParseForm()
	
	w.WriteHeader(http.StatusOK)

	//time
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	w.Header().Set("Server", "Servidor_Cecilia")
	w.Header().Set("Content-type", "text/html")

	t, _ := template.ParseFiles("static/index.html")
	
	if r.Method == "GET"{
				
		num1s := r.Form["num1"]
		num2s := r.Form["num2"]
		if len(num1s)==1 && len(num2s) ==1 {
			num1, err1 := strconv.ParseFloat(num1s[0], 64)
			num2, err2 := strconv.ParseFloat(num2s[0], 64)

			if err1 == nil && err2 == nil{
				soma = num1 + num2
			} else {
				error = "Argumento(s) inválidos"
			}
		}
	}
	
	a := Add_result{Soma: soma, Error: error}
	t.Execute(w, a)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.WriteHeader(http.StatusOK)

	//time, used to cache
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

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
	dispatch["somar"]   = adder

	server.ListenAndServe()
}
