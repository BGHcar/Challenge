package Controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"backend/Models" // Asegúrate de importar el paquete correcto donde se encuentra la estructura Email

	"github.com/joho/godotenv"
)

// carga las variables de entorno del archivo .env
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
}

// SearchEmail realiza una búsqueda de emails
func SearchEmail(w http.ResponseWriter, r *http.Request) {
	// Decodificar el JSON del cuerpo de la solicitud
	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	// Obtener el valor del parámetro 'term'
	term, ok := requestData["term"]
	if !ok || term == "" {
		http.Error(w, "El parámetro 'term' es requerido", http.StatusBadRequest)
		return
	}

	// Cargar variables de entorno
	loadEnv()

	// Leer las variables de entorno
	apiURL := os.Getenv("API_URL")
	zincUser := os.Getenv("ZINC_USER")
	zincPassword := os.Getenv("ZINC_PASSWORD")

	query := fmt.Sprintf(`{
        "search_type": "match",
        "query":
        {
            "term": "%s"
        },
        "from": 0,
        "max_results": 20,
        "_source": [
            "From", "To", "Subject", "body"
        ]
    }`, term)

	req, err := http.NewRequest("POST", apiURL+"EmailData/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(zincUser, zincPassword)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Leer la respuesta del servidor remoto
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		http.Error(w, "Error al decodificar la respuesta del servidor", http.StatusInternalServerError)
		return
	}

	// Extraer los correos electrónicos de la respuesta
	emails := make([]Models.Email, 0)
	hits := responseData["hits"].(map[string]interface{})
	hitsArray := hits["hits"].([]interface{})
	for _, hit := range hitsArray {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})
		from := source["From"].(string)
		to := source["To"].(string)
		subject := source["Subject"].(string)
		body := source["body"].(string)

		// Crear una instancia de Email y agregarla a la lista de correos electrónicos
		email := Models.Email{
			From:    from,
			To:      to,
			Subject: subject,
			Body:    body,
		}
		emails = append(emails, email)
	}

	// Codificar la lista de correos electrónicos como JSON y escribir en el cuerpo de la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}
