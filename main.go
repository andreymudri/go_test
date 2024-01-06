package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"./src/helper"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var savedLists []*ListNode

func saveListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// limpar valor previo de savedLists
	savedLists = nil
	// Decodificar o corpo da requisição JSON para uma estrutura de dados
	var requestData struct {
		List1 []int
		List2 []int
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	sort.Ints(requestData.List1)
	sort.Ints(requestData.List2)
	
	// transformar listas em listas encadeada
	list1 := &ListNode{Val: requestData.List1[0], Next: nil}
	for i := 1; i < len(requestData.List1); i++ {
		list1 = &ListNode{Val: requestData.List1[i], Next: list1}

	}
	list2 := &ListNode{Val: requestData.List2[0], Next: nil}
	for i := 1; i < len(requestData.List2); i++ {
		list2 = &ListNode{Val: requestData.List2[i], Next: list2}
	}

	// Salvar as listas recebidas em uma variável global
	savedLists = []*ListNode{list1, list2}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
}

// Função que recebe duas listas ligadas e retorna uma lista ligada mesclada
func mergeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	//checar se as listas estão vazias
	if savedLists[0] == nil || savedLists[1] == nil {
		w.Header().Set("Content-Type", "application/json")
	  w.WriteHeader(http.StatusNoContent)
	  json.NewEncoder(w).Encode("Listas vazias")
	  return
	}

	merged := mergeTwoLists(savedLists[0], savedLists[1])
	if merged == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// transformar lista ligada em lista de inteiros
	var mergedList []int
	for list := merged; list != nil; list = list.Next {
		mergedList = append(mergedList, list.Val)
	}
	// ordenar lista de inteiros
	sort.Ints(mergedList)

	// enviar a response merged
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mergedList)
}

// Função que recebe duas listas ligadas e retorna uma lista ligada numericamente crescente
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	mergedList := &ListNode{}
	head := mergedList
	// Percorrer as duas listas ligadas ao mesmo tempo
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	//adicionar o resto dos elementos da lista que não foi percorrida
	if list1 != nil {
		head.Next = list1
	} else if list2 != nil {
		head.Next = list2
	}

	return mergedList.Next
}
 func Helloworld(w http.ResponseWriter, r *http.Request) {
 	fmt.Fprintf(w, "Hello, World!")
 }

func main() {

	http.HandleFunc("/", Helloworld)

	http.HandleFunc("/saveLists", saveListHandler)

	http.HandleFunc("/merge", mergeHandler)

	http.ListenAndServe(":8080", nil)
}
