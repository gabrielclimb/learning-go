package router

import (
	"api-rest/controllers"
	"api-rest/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func HandleRequest() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"https://*", "http://*"}}))
	router.Use(middleware.ContentType)
	home(router)
	handlePersonalities(router)
	http.ListenAndServe(":8000", router)

}

func home(router *chi.Mux) {
	router.HandleFunc("/", controllers.Home)
}

func handlePersonalities(router *chi.Mux) {
	router.Route("/api/personalities", func(r chi.Router) {
		r.Get("/", controllers.AllPersonalities)
		r.Get("/{id}", controllers.GetPersonalityByID)
		r.Post("/", controllers.AddNewPersonality)
		r.Delete("/{id}", controllers.DeletePersonality)
		r.Put("/{id}", controllers.EditPersonality)
	})
}
