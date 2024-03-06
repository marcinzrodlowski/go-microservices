package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/marcinzrodlowski/go-microservices/handler"
	"net/http"
)

// loadRoutes returns a pointer to a chi.Mux instance (lightweight, fast router)
func loadRoutes() *chi.Mux {
	router := chi.NewRouter()     // create new chi router instance
	router.Use(middleware.Logger) // use middleware for logging (intercepts HTTP requests/responses)
	// define a GET route for the root URL
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes) // routes will be prefixed with /orders

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}
