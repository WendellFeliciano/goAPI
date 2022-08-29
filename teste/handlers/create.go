package handlers

import (
	"encoding/json"
	"fmt"
	"go_api/teste/models"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao decodificar o json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Erro ao inserir tarefa: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Tarefa inserida com sucesso! id: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
