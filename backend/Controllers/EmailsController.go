package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"

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
	term := requestData["term"]

	// Cargar variables de entorno
	loadEnv()

	// Leer las variables de entorno
	apiURL := os.Getenv("API_URL")
	zincUser := os.Getenv("ZINC_USER")
	zincPassword := os.Getenv("ZINC_PASSWORD")
	indexName := os.Getenv("INDEX_NAME")

	query := fmt.Sprintf(`{
        "search_type": "querystring",
        "query":
        {
            "term": "%s"
        },
        "from": 0,
        "max_results": 20,
        "_source": [
            "_id", "From", "To", "Subject","Metadata", "Message"
        ]
    }`, term)

	req, err := http.NewRequest("POST", apiURL+indexName+"/_search", strings.NewReader(query))
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
		_id := hitMap["_id"].(string)
		from := source["From"].(string)
		to := source["To"].(string)
		subject := source["Subject"].(string)
		Metadata := source["Metadata"].(map[string]interface{})
		Message := source["Message"].(string)

		// Crear una instancia de Email y agregarla a la lista de correos electrónicos
		email := Models.Email{
			Id:      _id,
			From:    from,
			To:      to,
			Subject: subject,
			Metadata: Models.Metadata{
				MimeVersion:             Metadata["Mime-Version"].(string),
				ContentType:             Metadata["Content-Type"].(string),
				ContentTransferEncoding: Metadata["Content-Transfer-Encoding"].(string),
				XFrom:                   Metadata["X-From"].(string),
				XTo:                     Metadata["X-To"].(string),
				XCc:                     Metadata["X-cc"].(string),
				XBcc:                    Metadata["X-bcc"].(string),
				XFolder:                 Metadata["X-Folder"].(string),
				XOrigin:                 Metadata["X-Origin"].(string),
				XFileName:               Metadata["X-FileName"].(string),
			},
			
			Message:    Message,
		}
		emails = append(emails, email)
	}

	// Codificar la lista de correos electrónicos como JSON y escribir en el cuerpo de la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}

// SearchEmail realiza una búsqueda de emails
func SearchAllEmails(w http.ResponseWriter, r *http.Request) {
	// Cargar variables de entorno
	loadEnv()

	// Leer las variables de entorno
	apiURL := os.Getenv("API_URL")
	zincUser := os.Getenv("ZINC_USER")
	zincPassword := os.Getenv("ZINC_PASSWORD")
	indexName := os.Getenv("INDEX_NAME")

	query := fmt.Sprintf(`{
        "search_type": "alldocuments",
        "from": 0,
        "max_results": 20,
        "_source": [
            "_id", "From", "To", "Subject","Metadata", "Message"
        ]
    }`)

	req, err := http.NewRequest("POST", apiURL+indexName+"/_search", strings.NewReader(query))
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
		_id := hitMap["_id"].(string)
		from := source["From"].(string)
		to := source["To"].(string)
		subject := source["Subject"].(string)
		Metadata := source["Metadata"].(map[string]interface{})
		Message := source["Message"].(string)

		// Crear una instancia de Email y agregarla a la lista de correos electrónicos
		email := Models.Email{
			Id:      _id,
			From:    from,
			To:      to,
			Subject: subject,
			Metadata: Models.Metadata{
				MimeVersion:             Metadata["Mime-Version"].(string),
				ContentType:             Metadata["Content-Type"].(string),
				ContentTransferEncoding: Metadata["Content-Transfer-Encoding"].(string),
				XFrom:                   Metadata["X-From"].(string),
				XTo:                     Metadata["X-To"].(string),
				XCc:                     Metadata["X-cc"].(string),
				XBcc:                    Metadata["X-bcc"].(string),
				XFolder:                 Metadata["X-Folder"].(string),
				XOrigin:                 Metadata["X-Origin"].(string),
				XFileName:               Metadata["X-FileName"].(string),
			},
			Message:    Message,
		}
		emails = append(emails, email)
	}

	// Codificar la lista de correos electrónicos como JSON y escribir en el cuerpo de la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}

func DeleteEmail(w http.ResponseWriter, r *http.Request) {
	// Código para eliminar un correo electrónico

	//Obtener el id del correo a eliminar
	id := chi.URLParam(r, "id")

	//Imprimir el id del correo a eliminar
	fmt.Println(id)

	// Cargar variables de entorno

	loadEnv()

	// Leer las variables de entorno

	apiURL := os.Getenv("API_URL")
	zincUser := os.Getenv("ZINC_USER")
	zincPassword := os.Getenv("ZINC_PASSWORD")
	indexName := os.Getenv("INDEX_NAME")

	// Crear la solicitud HTTP DELETE

	req, err := http.NewRequest("DELETE", apiURL+indexName+"/_doc/"+id, nil)

	if err != nil {
		log.Fatal(err)
	}

	// Establecer la autenticación básica en la solicitud HTTP

	req.SetBasicAuth(zincUser, zincPassword)

	// Realizar la solicitud HTTP

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Leer la respuesta HTTP

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Imprimir el código de estado y el cuerpo de la respuesta HTTP

	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))

	// Escribir el código de estado en la respuesta HTTP

	w.WriteHeader(resp.StatusCode)

	// Escribir el cuerpo de la respuesta HTTP

	w.Write(body)
}
