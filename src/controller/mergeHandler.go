package controller

import (
	"encoding/json"
	"net/http"
	"github.com/andreymudri/go_test/src/helper"
	"github.com/andreymudri/go_test/src/types"
)

// Função que recebe duas listas ligadas e retorna uma lista ligada mesclada
func MergeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	//checar se as listas estão vazias
	if SavedLists == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Listas vazias")
		return
	}
	if SavedLists[0] == nil || SavedLists[1] == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Lista vazias")
		return
	}
	merged := mergeTwoLists(SavedLists[0], SavedLists[1])
	if merged == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// transformar lista ligada em lista de inteiros
	result := helper.ListNodetoSortedArray(merged)
	// enviar a response merged
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Função que recebe duas listas ligadas e retorna uma lista ligada numericamente crescente
func mergeTwoLists(list1 *types.ListNode, list2 *types.ListNode) *types.ListNode {

	mergedList := &types.ListNode{}
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
	//limpa o savedLists
	SavedLists = nil

	return mergedList.Next
}
