package Controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi"

	"backend/Models" // Asegúrate de importar el paquete correcto donde se encuentra la estructura Email

	"github.com/joho/godotenv"
)

// loadEnv carga las variables de entorno del archivo .env
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

	page := chi.URLParam(r, "page")

	// Obtener el valor del parámetro 'term'
	term := requestData["term"]

	pageInt, err:= strconv.Atoi(page)	
	if err != nil {
		http.Error(w, "Error al convertir el número de página", http.StatusInternalServerError)
		return
	}

	if pageInt > 0 {
		pageInt--
		fmt.Print("Esta es la pagina : ", pageInt)
		pageInt = pageInt * 20

	}	
	page = strconv.Itoa(pageInt)


	// Cargar variables de entorno
	loadEnv()

	// Leer las variables de entorno
	apiURL := os.Getenv("API_URL")
	zincUser := os.Getenv("ZINC_USER")
	zincPassword := os.Getenv("ZINC_PASSWORD")
	indexName := os.Getenv("INDEX_NAME")

	query := fmt.Sprintf(`{
        "search_type": "querystring",
        "query": {
            "term": "%s"
        },
        "from": %s,
        "max_results": 20,
        "_source": [
            "_id", "From", "To", "Subject", "Metadata", "Message"
        ]
    }`, term, page)

	fmt.Println("Este es el query : ", query)

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

	// Verificar si el campo "hits" está presente en la respuesta
	hitsData, ok := responseData["hits"].(map[string]interface{})
	if !ok || hitsData == nil {
		http.Error(w, "El campo 'hits' en la respuesta está vacío o no está presente", http.StatusInternalServerError)
		return
	}

	// Obtener el total de correos electrónicos
	totalHits, ok := hitsData["total"].(map[string]interface{})
	if !ok {
		http.Error(w, "Error al obtener el número total de correos electrónicos", http.StatusInternalServerError)
		return
	}

	// Extraer el valor de "value" si existe
	value, ok := totalHits["value"].(float64)
	if !ok {
		http.Error(w, "Error al convertir el número total de correos electrónicos", http.StatusInternalServerError)
		return
	}

	emails := Models.Emails{}
	emails.Total = int(value)

	hitsArray := hitsData["hits"].([]interface{})
	for _, hit := range hitsArray {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})
		_id := hitMap["_id"].(string)
		from := source["From"].(string)
		to := source["To"].(string)
		subject := source["Subject"].(string)
		Metadata := source["Metadata"].(map[string]interface{})
		Message := source["Message"].(string)

		email := Models.Email{
			Id:   _id,
			From: from,
			To:   to,
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
			Message: Message,
		}
		emails.Emails = append(emails.Emails, email)
	}

	// Codificar la lista de correos electrónicos como JSON y escribir en el cuerpo de la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}

// SearchAllEmails realiza una búsqueda de todos los correos electrónicos
func SearchAllEmails(w http.ResponseWriter, r *http.Request) {
	// Cargar variables de entorno
	loadEnv()

	// Leer las variables de entorno
	apiURL := os.Getenv("API_URL")
	zincUser := os.Getenv("ZINC_USER")
	zincPassword := os.Getenv("ZINC_PASSWORD")
	indexName := os.Getenv("INDEX_NAME")

	page := chi.URLParam(r, "page")


	pageInt, err:= strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Error al convertir el número de página", http.StatusInternalServerError)
		return
	}

	if pageInt > 0 {
		pageInt--
		fmt.Print("Esta es la pagina : ", pageInt)
		pageInt = pageInt * 20

	}

	page = strconv.Itoa(pageInt)
	


	query := fmt.Sprintf(`{
        "search_type": "alldocuments",
        "from": %s,
        "max_results": 20,
        "_source": [
            "_id", "From", "To", "Subject", "Metadata", "Message"
        ]
    }`, page)

	fmt.Println("Este es el query : ", query)

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

	// Verificar si el campo "hits" está presente en la respuesta
	hitsData, ok := responseData["hits"].(map[string]interface{})
	if !ok || hitsData == nil {
		http.Error(w, "El campo 'hits' en la respuesta está vacío o no está presente", http.StatusInternalServerError)
		return
	}

	// Obtener el total de correos electrónicos
	totalHits, ok := hitsData["total"].(map[string]interface{})
	if !ok {
		http.Error(w, "Error al obtener el número total de correos electrónicos", http.StatusInternalServerError)
		return
	}

	// Extraer el valor de "value" si existe
	value, ok := totalHits["value"].(float64)
	if !ok {
		http.Error(w, "Error al convertir el número total de correos electrónicos", http.StatusInternalServerError)
		return
	}

	emails := Models.Emails{}
	emails.Total = int(value)

	hitsArray := hitsData["hits"].([]interface{})
	for _, hit := range hitsArray {
		hitMap := hit.(map[string]interface{})
		source := hitMap["_source"].(map[string]interface{})
		_id := hitMap["_id"].(string)
		from := source["From"].(string)
		to := source["To"].(string)
		subject := source["Subject"].(string)
		Metadata := source["Metadata"].(map[string]interface{})
		Message := source["Message"].(string)

		email := Models.Email{
			Id:   _id,
			From: from,
			To:   to,
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
			Message: Message,
		}
		emails.Emails = append(emails.Emails, email)
	}

	// Codificar la lista de correos electrónicos como JSON y escribir en el cuerpo de la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}

// DeleteEmail elimina un correo electrónico por su ID
func DeleteEmail(w http.ResponseWriter, r *http.Request) {
	//Obtener el id del correo a eliminar
	id := chi.URLParam(r, "id")

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
