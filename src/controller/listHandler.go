package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andreymudri/go_test/src/helper"
	"github.com/andreymudri/go_test/src/types"
)
var SavedLists []*types.ListNode

func SaveListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// limpar valor previo de savedLists
	SavedLists = nil
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

	// transformar listas em listas encadeada
	list1 := helper.CreateNodeList(requestData.List1)
	list2 := helper.CreateNodeList(requestData.List2)

	// Salvar as listas recebidas em uma variável global
SavedLists = []*types.ListNode{list1, list2}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
}


