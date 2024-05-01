package main

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi"
    "github.com/go-chi/cors"
    "backend/routes"
)

func main() {
    // Crear un enrutador Chi
    r := chi.NewRouter()

    // Configurar el middleware CORS
    cors := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:8081"}, // Permitir solicitudes desde este origen
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Métodos permitidos
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // Cabeceras permitidas
        AllowCredentials: true, // Permitir enviar credenciales (cookies)
        MaxAge:           300,  // Duración de la caché de preflight (en segundos)
    })

    // Usar el middleware CORS en todas las rutas
    r.Use(cors.Handler)

    // Establecer las rutas de los emails
    routes.SetEmailRoutes(r)

    // Iniciar el servidor
    fmt.Println("Servidor iniciado en el puerto : http://localhost:9000")
    http.ListenAndServe(":9000", r)
}
