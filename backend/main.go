package main

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi"
    "backend/routes"
)

func main() {
    // Crear un enrutador Chi
    r := chi.NewRouter()

    // Establecer las rutas de los emails
    routes.SetEmailRoutes(r)

    // Iniciar el servidor
    fmt.Println("Servidor iniciado en el puerto : http://localhost:9000")
    http.ListenAndServe(":9000", r)
}
