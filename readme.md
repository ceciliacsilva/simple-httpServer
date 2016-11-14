HTTP-Server
-------------------

"A web server is a computer system that processes requests via HTTP, the basic network protocol used to distribute information on the World Wide Web. The term can refer to the entire system, or specifically to the software that accepts and supervises the HTTP requests."[1]

GO Language
--------------------

"Go (often referred to as golang) is a free and open source programming language created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing, memory safety features and CSP-style concurrent programming features added."[2]

Use
--------------------

	$ go run server.go

Open a browser:

        http://localhost:8080/


Programs
-------------------

Statics - Files

	- index.html
	- somar.html
	- texto.html

Server

	- server.go

Function
-------------------

- Adder two numbers: adder, show an error line if the arguments are invalids (200)
- Moved Permanently: redirect to URL-Location (404)
- Text in different languages: show "hello World" in portuguese, france or english using "Accept-Language" field of header (200)

Reference
--------------------

[1] https://en.wikipedia.org/wiki/Web_server

[2] https://en.wikipedia.org/wiki/Go_(programming_language)

Other Links
--------------------

http://thenewstack.io/building-a-web-server-in-go/

https://tutorialedge.net/creating-a-simple-web-server-with-go-golang

https://golang.org/src/net/http/request.go?h=Header%28%29.Set

https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/03.2.html

http://stackoverflow.com/questions/16896430/golang-http-server-cant-get-post-value

https://github.com/WanZheng/golang-simple-http-server/blob/master/main.go

https://github.com/kisielk/jsonrpc-example/blob/master/server.go#L28

https://www.nicolasmerouze.com/middlewares-golang-best-practices-examples/