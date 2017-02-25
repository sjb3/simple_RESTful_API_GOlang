package main

import (
	"io"
	"net/http"
	"fmt"
)

//type DogHandler int

func dogDog(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "bow wow wow")
}

//type CatHandler int

func catCat(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "meow, meow")
}

func main() {
	// var dog DogHandler
	// var cat CatHandler
	//
	// mux := http.NewServeMux()
	// mux.Handle("/dog", dog)
	// mux.Handle("/cat", cat)
	http.HandleFunc("/", dogDog)
	http.HandleFunc("/cat/", catCat)

	fmt.Println("server running")
	// http.ListenAndServe(":9000", mux)
	http.ListenAndServe(":9000", nil)
}

//Refactor above using HandleFunc()

//func main() {
//
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
//		io.WriteString(res, "cha cha cha")
//	})
//
//	mux.HandleFunc("/cat/", func(res http.ResponseWriter, req *http.Request) {
//		io.WriteString(res, "yum yum yum")
//	})
//	fmt.Println("server running")
//	http.ListenAndServe(":9000", mux)
//}
