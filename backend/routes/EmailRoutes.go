package routes

import (
    "backend/Controllers"
    "github.com/go-chi/chi"
)

// SetEmailRoutes establece las rutas relacionadas con los emails
func SetEmailRoutes(r *chi.Mux) {
    r.Post("/search", Controllers.SearchEmail)
}
