package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HnadleFunc("/", func(w.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello Docker!")
	})

	http.ListenAndServe(":8080", nil)
}