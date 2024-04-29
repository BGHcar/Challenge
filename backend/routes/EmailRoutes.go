package routes

import (
    "backend/Controllers"
    "github.com/go-chi/chi"
)

// SetEmailRoutes establece las rutas relacionadas con los emails
func SetEmailRoutes(r *chi.Mux) {
    r.Post("/search", Controllers.SearchEmail) // Fixed function call
    r.Delete("/delete/{id}", Controllers.DeleteEmail) // Added missing function call
    r.Post("/searchall", Controllers.SearchAllEmails) // Added missing function call
}
