package main

import (
	"fmt"
	"net/http"

	"github.com/andreymudri/go_test/src/controller"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {

	http.HandleFunc("/", Helloworld)

	http.HandleFunc("/saveLists", controller.SaveListHandler)

	http.HandleFunc("/merge", controller.MergeHandler)

	http.ListenAndServe(":8080", nil)
}
