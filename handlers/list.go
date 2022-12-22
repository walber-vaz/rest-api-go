package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error ao obter registro: %v", err)
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(todos)
}
