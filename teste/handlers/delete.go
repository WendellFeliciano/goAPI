package handlers

import (
	"encoding/json"
	"go_api/teste/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	row, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if row > 1 {
		log.Printf("Erro: foram removidos %d registros.", row)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "registro removido com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
