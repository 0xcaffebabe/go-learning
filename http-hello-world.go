package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/1", func(writer http.ResponseWriter,request *http.Request){
		writer.Write([]byte("hello, " + request.URL.RawQuery))
	})
	http.HandleFunc("/2", func(writer http.ResponseWriter,request *http.Request){
		writer.Write([]byte("word"))
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}