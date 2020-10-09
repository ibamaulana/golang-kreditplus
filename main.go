package main

import (
	"github.com/ibamaulana/golang-kreditplus/controller/front"
)

import "fmt"
import "net/http"

func main() {
	http.Handle("/static/", 
	http.StripPrefix("/static/", 
		http.FileServer(http.Dir("assets"))))
		
	http.HandleFunc("/", front.GetController)

    fmt.Println("server started at localhost:9000")
    http.ListenAndServe(":9000", nil)
}
