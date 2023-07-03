package router

import (
	"api-rest/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleRequest() {
	router := chi.NewRouter()
	home(router)
	handlePersonalities(router)
	http.ListenAndServe(":8000", router)
}

func home(router *chi.Mux) {
	router.HandleFunc("/", controllers.Home)
}

func handlePersonalities(router *chi.Mux) {
	router.Route("/personalities", func(r chi.Router) {
		r.Get("/", controllers.AllPersonalities)
		r.Get("/{id}", controllers.GetPersonalityByID)
	})
}
