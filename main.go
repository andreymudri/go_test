package main

import (
	"fmt"
	"net/http"
)
type ListNode struct {
	Val  int
	Next *ListNode
}


func saveListHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement code for function
}
func mergeHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement code for function
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// TODO: escreva seu código
	return nil
}


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	
	http.HandleFunc("/saveLists", saveListHandler)
	http.HandleFunc("/merge", mergeHandler)
		
	http.ListenAndServe(":8080", nil)
}