package controllers

import (
	"api-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func GetPersonalityByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(fmt.Sprintf("ID:%s", id))
	for _, personalities := range models.Personalities {
		if strconv.Itoa(personalities.ID) == id {
			json.NewEncoder(w).Encode(personalities)
		}

	}
}
