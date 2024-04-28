package main

import (

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

)


func Routes() *chi.Mux{
	mux := chi.NewMux()


	// global middlewares
	mux.Use(
		middleware.Logger,  	  // Log API request calls
		middleware.Recoverer,  // Recover from panics without crashing server
	)

	mux.Post(	"/es/EmailData/_search", 	nil)

	return mux
}