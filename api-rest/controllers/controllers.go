package controllers

import (
	"api-rest/database"
	"api-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.First(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	id := chi.URLParam(r, "id")
	fmt.Println(fmt.Sprintf("ID:%s", id))
	database.DB.Find(&personality, id)
	json.NewEncoder(w).Encode(personality)
}
