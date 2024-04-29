package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "net/http/pprof"  // Importar el paquete pprof
)

var count int = 0 // Contador de correos procesados

// Definición de la estructura Metadata
type Metadata struct {
	MimeVersion             string `json:"Mime-Version"`
	ContentType             string `json:"Content-Type"`
	ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
	XFrom                   string `json:"X-From"`
	XTo                     string `json:"X-To"`
	XCc                     string `json:"X-cc"`
	XBcc                    string `json:"X-bcc"`
	XFolder                 string `json:"X-Folder"`
	XOrigin                 string `json:"X-Origin"`
	XFileName               string `json:"X-FileName"`
}

// Definición de la estructura Email
type Email struct {
	MessageID string   `json:"Message-ID"`
	Date      string   `json:"Date"`
	From      string   `json:"From"`
	To        string   `json:"To"`
	Subject   string   `json:"Subject"`
	Metadata  Metadata `json:"Metadata"`
	Body      string   `json:"body"`
}

// Definición de la estructura principal para el correo electrónico
type EmailData struct {
	Index   string `json:"index"`
	Records []Email  `json:"records"`
}

func main() {
	start := time.Now() // Iniciar el temporizador

	// Cargar las variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error cargando el archivo .env:", err)
		return
	}

	// Obtener la ruta de la carpeta desde la variable de entorno
	ruta := os.Getenv("PATH_DIRECTION")
	if ruta == "" {
		fmt.Println("Variable de entorno PATH_DIRECTION no encontrada en el archivo .env")
		return
	}

	fmt.Println("Procesando correos electrónicos en la carpeta:", ruta)

	// Iniciar el servidor HTTP para pprof en segundo plano para el perfilado de la aplicación con pprof en http://localhost:6060/debug/pprof/
	go func() {
		fmt.Println("Perfilado de pprof está disponible en http://localhost:6060/debug/pprof/")
		http.ListenAndServe("localhost:6060", nil)
	}()

	// Recorrer de manera recursiva la carpeta y leer los archivos
	err := filepath.Walk(ruta, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Verificar si el archivo no es un directorio y no tiene extensión
		if !info.IsDir() && filepath.Ext(info.Name()) == "" {
			// Leer el contenido del archivo
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Procesar el contenido del archivo como correo electrónico
			email := procesarCorreo(string(content))

			// Incrementar el contador de correos procesados
			count++

			// Crear un objeto EmailData con el correo electrónico obtenido
			emailData := EmailData{
				Index:   "EmailData",
				Records: []Email{email}, // Aquí cambia email a []Email{email}
			}

			// Convertir el objeto EmailData a formato JSON
			jsonData, err := json.MarshalIndent(emailData, "", "    ")
			if err != nil {
				fmt.Println("Error al convertir a JSON:", err)
				return err
			}

			sendDataToZincSearch(jsonData) // Enviar el correo electrónico a ZincSearch
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error al leer los archivos:", err)
		return
	}

	elapsed := time.Since(start) // Calcular el tiempo transcurrido
	fmt.Printf("Tiempo total de indexación: %s\n", elapsed) // Imprimir el tiempo total de indexación
	fmt.Printf("Total de emails procesados : %d\n", count)  // Imprimir el total de correos procesados
}

// Función para procesar el contenido del archivo como un correo electrónico
func procesarCorreo(content string) Email {
	// Dividir el contenido del correo por líneas
	lines := strings.Split(content, "\n")

	// Crear un nuevo objeto Email y asignar los valores de los campos
	email := Email{}
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			switch key {
			case "Message-ID":
				email.MessageID = fmt.Sprintf(value)
			case "Date":
				email.Date = fmt.Sprintf(value)
			case "From":
				email.From = fmt.Sprintf(value)
			case "To":
				email.To = fmt.Sprintf(value)
			case "Subject":
				email.Subject = fmt.Sprintf(value)
			case "Mime-Version":
				email.Metadata.MimeVersion = fmt.Sprintf(value)
			case "Content-Type":
				email.Metadata.ContentType = fmt.Sprintf(value)
			case "Content-Transfer-Encoding":
				email.Metadata.ContentTransferEncoding = fmt.Sprintf(value)
			case "X-From":
				email.Metadata.XFrom = fmt.Sprintf(value)
			case "X-To":
				email.Metadata.XTo = fmt.Sprintf(value)
			case "X-cc":
				email.Metadata.XCc = fmt.Sprintf(value)
			case "X-bcc":
				email.Metadata.XBcc = fmt.Sprintf(value)
			case "X-Folder":
				email.Metadata.XFolder = fmt.Sprintf(value)
			case "X-Origin":
				email.Metadata.XOrigin = fmt.Sprintf(value)
			case "X-FileName":
				email.Metadata.XFileName = fmt.Sprintf(value)
			default:
				// Ignorar otras líneas que no corresponden a los campos deseados
			}
		} else if len(parts) == 1 && parts[0] == "" {
			// Una línea en blanco indica el fin de los metadatos y el comienzo del cuerpo del correo
			break
		}
	}

	// El cuerpo del correo es el contenido restante
	email.Body = fmt.Sprintf(strings.Join(lines, "\n"))

	return email
}

// Función para enviar el correo electrónico a ZincSearch
func sendDataToZincSearch(email []byte) {
	apiURL := os.Getenv("API_URL")
	username := os.Getenv("ZINC_USER")
	password := os.Getenv("ZINC_PASSWORD")

	// Crear una solicitud HTTP POST con el JSON de emailData
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(email))
	if err != nil {
		fmt.Println("Error al crear la solicitud HTTP:", err)
		return
	}

	// Establecer la autenticación básica en la solicitud HTTP
	req.SetBasicAuth(username, password)

	// Establecer el encabezado Content-Type en la solicitud HTTP
	req.Header.Set("Content-Type", "application/json")

	// Realizar la solicitud HTTP
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta HTTP
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta HTTP:", err)
		return
	}

	// Imprimir el código de estado y el cuerpo de la respuesta HTTP
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}
