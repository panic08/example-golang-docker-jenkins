package main

import (
	"fmt"
	"net/http"
)

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hi from e %s!", r.URL.Path[1:])
	http.ServeFile(response, request, "Index.html")
}

func main() {
	fmt.Println("Server Starting")
	http.HandleFunc("/", HttpFileHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//http.HandleFunc("/", indexTemplateHandler)

	http.ListenAndServe(":8080", nil)
}
