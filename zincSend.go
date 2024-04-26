package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Leer el contenido de output.json
	file, err := os.Open("output.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Leer el contenido de output.json
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Crear la solicitud HTTP para ZincSearch
	req, err := http.NewRequest("POST", os.Getenv("ZINC_SEARCH_API_PATH"), strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	// Configurar la autenticación básica si es necesario
	req.SetBasicAuth(os.Getenv("ZINC_SEARCH_API_USER"), os.Getenv("ZINC_SEARCH_API_PASS"))

	// Establecer el encabezado Content-Type
	req.Header.Set("Content-Type", "application/json")

	// Realizar la solicitud HTTP
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	log.Println(resp.StatusCode)

	// Leer y mostrar la respuesta del servidor
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
