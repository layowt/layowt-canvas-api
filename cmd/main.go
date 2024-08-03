package main

import (
	"fmt"
	"net/http"

	"github.com/layowt/layowt-canvas-api/types"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	user := types.User{
		Uid: 1,
	}

	fmt.Println(user)
}
