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
	database.DB.Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	id := chi.URLParam(r, "id")
	fmt.Println(fmt.Sprintf("ID:%s", id))
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func AddNewPersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonality models.Personality
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&newPersonality)
	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(&newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality

	id := chi.URLParam(r, "id")
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(&personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	id := chi.URLParam(r, "id")
	database.DB.First(&personality, id)

	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)
}
