package main

import (
	"net/http"
)

func indexHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a test"))
}

func main() {
	http.HandleFunc("/", indexHander)
	http.ListenAndServe(":8080", nil)
}
