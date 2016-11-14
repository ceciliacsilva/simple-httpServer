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
        "net/http"
	"time"
	"html/template"
	"strconv"
	"regexp"
	"log"
)

//like hash table
var dispatch map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}
type Add_result struct {
	Soma   float64
	Error  string
}

type text_pag struct{
	Title  string
	Enc    string
	Text   string
	Button string
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

	//time, used to cache
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	//name server and content-type
	w.Header().Set("Server", "Servidor-Cecilia")
	w.Header().Set("Content-type", "text/html")
	
	//status-line
	w.WriteHeader(http.StatusOK)

	t, _ := template.ParseFiles("static/index.html")

	t.Execute(w, nil)
}

func adder(w http.ResponseWriter, r *http.Request){
	var soma  float64 = 0.0
	var error string  = " " 
	r.ParseForm()
	
	//time
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	w.Header().Set("Server", "Servidor_Cecilia")
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	
	t, _ := template.ParseFiles("static/somar.html")
	
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

func moved(w http.ResponseWriter, r *http.Request){
	
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	w.Header().Set("Server", "Servidor_Cecilia")
	w.Header().Set("Content-type", "text/html")

	//location moved
	w.Header().Set("Location", "http://www.lrc.eletrica.ufu.br/")

	w.WriteHeader(http.StatusMovedPermanently)
}

func texto(w http.ResponseWriter, r *http.Request){
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	//name server and content-type
	w.Header().Set("Server", "Servidor-Cecilia")
	w.Header().Set("Content-type", "text/html")
	
	w.WriteHeader(http.StatusOK)

	t, _ := template.ParseFiles("static/texto.html")

	france      := text_pag{Title: "Texte", Enc: "fr-FR", Text: "Bonjour le monde!", Button: "Retour"}
	english     := text_pag{Title: "Text",  Enc: "en-US", Text: "Hello World!",      Button: "Back"}
	portuguese  := text_pag{Title: "Texto", Enc: "pt-BR", Text: "Oi mundo!",         Button: "Voltar"}
		
	accept_lang := r.Header.Get("Accept-Language")
	re          := regexp.MustCompile("[a-z]*-[A-Z]*")
	lang        := re.FindAllString(accept_lang, -1)

	if len(lang) >= 1{
		language := lang[0]
		if language == "fr-FR"{
			t.Execute(w, france)
		} else if language == "pt-BR" {
			t.Execute(w, portuguese)
		} else {
			t.Execute(w, english)
		}
	}
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler:  &myHandler{},
	}

	dispatch = make(map[string]func(http.ResponseWriter, *http.Request))
	
	dispatch["somar"]   = adder
	dispatch["movido"]  = moved
	dispatch["texto"]   = texto
	
	err := server.ListenAndServe()

	if err!=nil{
		log.Fatal("ListenAndServe: ", err)
	}
}

//show files
//http.ServeFile(w, r, "static/index.html") 
