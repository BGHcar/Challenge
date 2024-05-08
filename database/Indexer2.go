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
	"github.com/pkg/profile"

	"github.com/joho/godotenv"
	_ "net/http/pprof" // Importar para habilitar el profiling mediante net/http/pprof
)

var (
	apiURL   string
	username string
	password string
	indexName string
	count int // Contador de correos procesados
)

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
	Message   string   `json:"Message"`
}

// Definición de la estructura principal para el correo electrónico
type EmailData struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// Iniciar el profiling
	go func() {
		fmt.Println("Profiling server listening on http://localhost:6060/debug/pprof/")
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

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

	indexName := os.Getenv("INDEX_NAME")
	apiURL := os.Getenv("API_URL")
	username := os.Getenv("ZINC_USER")
	password := os.Getenv("ZINC_PASSWORD")

	fmt.Println("Procesando correos electrónicos en la carpeta:", ruta)

	// Recorrer de manera recursiva la carpeta y leer los archivos
	var emailData EmailData
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
			email, err := procesarCorreo(string(content))
			if err != nil {
				fmt.Printf("Error al procesar el correo electrónico en el archivo %s: %v\n", path, err)
				return nil
			}

			// Incrementar el contador de correos procesados
			count++
			fmt.Printf("Procesando correo electrónico %d\n", count)

			// Agregar el correo electrónico al slice de registros
			emailData.Records = append(emailData.Records, email)

			// Si el número de correos procesados alcanza un límite (por ejemplo, 100), enviar los datos a la API
			if count%10000 == 0 {
				// Asignar el nombre del índice al objeto EmailData
				emailData.Index = indexName
				fmt.Printf("Procesando correo electrónico %d: \n", count)

				// Enviar los correos electrónicos procesados a la API
				if err := sendDataToZincSearch(emailData, apiURL, username, password); err != nil {
					fmt.Println("Error al enviar los correos electrónicos a ZincSearch:", err)
					return err
				}

				// Reiniciar emailData para prepararlo para los siguientes correos electrónicos
				emailData = EmailData{}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error al leer los archivos:", err)
		return
	}

	elapsed := time.Since(start)                            // Calcular el tiempo transcurrido
	fmt.Printf("Tiempo total de indexación: %s\n", elapsed) // Imprimir el tiempo total de indexación
	fmt.Printf("Total de emails procesados : %d\n", count)  // Imprimir el total de correos procesados

	// Si quedan correos electrónicos en emailData, enviarlos a la API
	if len(emailData.Records) > 0 {
		// Asignar el nombre del índice al objeto EmailData
		emailData.Index = indexName

		// Enviar los correos electrónicos procesados a la API
		if err := sendDataToZincSearch(emailData, apiURL, username, password); err != nil {
			fmt.Println("Error al enviar los correos electrónicos a ZincSearch:", err)
			return
		}
	}
}

// Función para procesar el contenido del archivo como un correo electrónico
func procesarCorreo(content string) (Email, error) {
	// Dividir el contenido del correo por líneas
	lines := strings.Split(content, "\n")

	// Crear un nuevo objeto Email y asignar los valores de los campos
	email := Email{}
	for i, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			switch key {
			case "Message-ID":
				email.MessageID = value
			case "Date":
				email.Date = value
			case "From":
				email.From = value
			case "To":
				email.To = value
			case "Subject":
				email.Subject = value
			case "Mime-Version":
				email.Metadata.MimeVersion = value
			case "Content-Type":
				email.Metadata.ContentType = value
			case "Content-Transfer-Encoding":
				email.Metadata.ContentTransferEncoding = value
			case "X-From":
				email.Metadata.XFrom = value
			case "X-To":
				email.Metadata.XTo = value
			case "X-cc":
				email.Metadata.XCc = value
			case "X-bcc":
				email.Metadata.XBcc = value
			case "X-Folder":
				email.Metadata.XFolder = value
			case "X-Origin":
				email.Metadata.XOrigin = value
			case "X-FileName":
				email.Metadata.XFileName = value
			default:
				// Ignorar otras líneas que no corresponden a los campos deseados
			}
		} else if len(parts) == 1 && parts[0] == "" {
			// Una línea en blanco indica el fin de los metadatos y el comienzo del cuerpo del correo
			email.Message = strings.Join(lines[i+1:], "\n")
			break
		}
	}

	return email, nil
}

// Función para enviar el correo electrónico a ZincSearch
func sendDataToZincSearch(emailData EmailData, apiURL, username, password string) error {
	// Convertir el objeto EmailData a formato JSON
	jsonData, err := json.MarshalIndent(emailData, "", "    ")
	if err != nil {
		return fmt.Errorf("error al convertir a JSON: %v", err)
	}

	// Crear una solicitud HTTP POST con el JSON de emailData
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error al crear la solicitud HTTP: %v", err)
	}

	// Establecer la autenticación básica en la solicitud HTTP
	req.SetBasicAuth(username, password)

	// Establecer el encabezado Content-Type en la solicitud HTTP
	req.Header.Set("Content-Type", "application/json")

	// Realizar la solicitud HTTP
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error al realizar la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta HTTP
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error al leer la respuesta HTTP: %v", err)
	}

	// Imprimir el código de estado y el cuerpo de la respuesta HTTP
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))

	return nil
}
