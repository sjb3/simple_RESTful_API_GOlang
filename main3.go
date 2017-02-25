package main

import (
	"io"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","text/html; charset=utf-8")

	switch req.URL.Path{
	case "/cat":
		io.WriteString(res, `<img src="https://www.petdrugsonline.co.uk/images/page-headers/cats-master-header" />`)

	case "/dog":
		io.WriteString(res, `<img src="http://cdn2-www.dogtime.com/assets/uploads/gallery/30-impossibly-cute-puppies/impossibly-cute-puppy-8.jpg" />`)
	}
}

func main(){
	var h myHandler
	http.ListenAndServe(":9000", h)
}
