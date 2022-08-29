package handlers

import (
	"encoding/json"
	"go_api/teste/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(todos)
}
