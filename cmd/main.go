package main

import (
	"fmt"
	"net/http"

	"github.com/layowt/layowt-canvas-api/types"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// map the '/hello-world' route to the helloWorldHandler function
	http.HandleFunc("/hello-world", helloWorldHandler)
	http.ListenAndServe(":8080", nil)

	user := types.User{
		Uid: 1,
	}

	admin := types.Admin.DeleteUser(nil)

	fmt.Println(user)
	fmt.Println(admin)
}
