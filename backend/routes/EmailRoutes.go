package routes

import (
    "backend/Controllers" // Fixed import path
    "github.com/go-chi/chi"
)

// SetEmailRoutes establece las rutas relacionadas con los emails
func SetEmailRoutes(r *chi.Mux) {
    r.Post("/search/{page}", Controllers.SearchEmail) // Fixed function call
    r.Delete("/delete/{id}", Controllers.DeleteEmail) // Added missing function call
    r.Get("/searchall/{page}", Controllers.SearchAllEmails) // Added missing function call
}
