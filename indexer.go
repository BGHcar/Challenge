package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type email struct {
	MessageID                  string `json:"Message-ID"`
	Date                       string `json:"Date"`
	From                       string `json:"From"`
	To                         string `json:"To"`
	Subject                    string `json:"Subject"`
	MimeVersion                string `json:"Mime-Version"`
	ContentType                string `json:"Content-Type"`
	ContentTransferEncoding    string `json:"Content-Transfer-Encoding"`
	XFrom                      string `json:"X-From"`
	XTo                        string `json:"X-To"`
	Xcc                        string `json:"X-cc"`
	Xbcc                       string `json:"X-bcc"`
	XFolder                    string `json:"X-Folder"`
	XOrigin                    string `json:"X-Origin"`
	XFileName                  string `json:"X-FileName"`
	Body                       string `json:"body"`
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	// Obtener la ruta del directorio maildir desde la variable de entorno
	rootDir := os.Getenv("PATH_DIRECTION")

	// Crear un archivo para guardar la salida en formato JSON
	outputFile, err := os.Create("output.json")
	if err != nil {
		log.Fatal("Error creando archivo de salida:", err)
	}
	defer outputFile.Close()

	// Crear un encoder JSON con las opciones adecuadas para evitar el escape de ciertos caracteres
	encoder := json.NewEncoder(outputFile)
	encoder.SetEscapeHTML(false)

	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Si el archivo es un directorio, no hacemos nada
		if info.IsDir() {
			return nil
		}

		// Si el archivo es un archivo plano, lo procesamos
		if filepath.Ext(path) == "" {
			wg.Add(1)
			go func(filePath string) {
				defer wg.Done()
				if emailData, err := indexFile(filePath); err != nil {
					fmt.Printf("Error al indexar archivo %s: %v\n", filePath, err)
				} else {
					if err := encoder.Encode(emailData); err != nil {
						fmt.Printf("Error al escribir en el archivo: %v\n", err)
					}
				}
			}(path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Tiempo total de indexación: %s\n", elapsed)
}

// Función para indexar un archivo de correo plano
func indexFile(filePath string) (*email, error) {
	// Verificar si el archivo es un acceso directo y saltarlo si es así
	if filepath.Ext(filePath) == ".lnk" {
		fmt.Printf("Saltando acceso directo: %s\n", filePath)
		return nil, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()


	const maxCapacity = 1024 * 1024 // 1 MB
	scanner := bufio.NewScanner(file)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)


	var emailData email
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Message-ID:") {
			emailData.MessageID = strings.TrimSpace(strings.TrimPrefix(line, "Message-ID:"))
		} else if strings.HasPrefix(line, "Date:") {
			emailData.Date = strings.TrimSpace(strings.TrimPrefix(line, "Date:"))
		} else if strings.HasPrefix(line, "From:") {
			emailData.From = strings.TrimSpace(strings.TrimPrefix(line, "From:"))
		} else if strings.HasPrefix(line, "To:") {
			emailData.To = strings.TrimSpace(strings.TrimPrefix(line, "To:"))
		} else if strings.HasPrefix(line, "Subject:") {
			emailData.Subject = strings.TrimSpace(strings.TrimPrefix(line, "Subject:"))
		} else if strings.HasPrefix(line, "Mime-Version:") {
			emailData.MimeVersion = strings.TrimSpace(strings.TrimPrefix(line, "Mime-Version:"))
		} else if strings.HasPrefix(line, "Content-Type:") {
			emailData.ContentType = strings.TrimSpace(strings.TrimPrefix(line, "Content-Type:"))
		} else if strings.HasPrefix(line, "Content-Transfer-Encoding:") {
			emailData.ContentTransferEncoding = strings.TrimSpace(strings.TrimPrefix(line, "Content-Transfer-Encoding:"))
		} else if strings.HasPrefix(line, "X-From:") {
			emailData.XFrom = strings.TrimSpace(strings.TrimPrefix(line, "X-From:"))
		} else if strings.HasPrefix(line, "X-To:") {
			emailData.XTo = strings.TrimSpace(strings.TrimPrefix(line, "X-To:"))
		} else if strings.HasPrefix(line, "X-cc:") {
			emailData.Xcc = strings.TrimSpace(strings.TrimPrefix(line, "X-cc:"))
		} else if strings.HasPrefix(line, "X-bcc:") {
			emailData.Xbcc = strings.TrimSpace(strings.TrimPrefix(line, "X-bcc:"))
		} else if strings.HasPrefix(line, "X-Folder:") {
			emailData.XFolder = strings.TrimSpace(strings.TrimPrefix(line, "X-Folder:"))
		} else if strings.HasPrefix(line, "X-Origin:") {
			emailData.XOrigin = strings.TrimSpace(strings.TrimPrefix(line, "X-Origin:"))
		} else if strings.HasPrefix(line, "X-FileName:") {
			emailData.XFileName = strings.TrimSpace(strings.TrimPrefix(line, "X-FileName:"))
		} else if line == "" {
			break // Fin del encabezado, comienza el cuerpo del mensaje
		}
	}

	// Leer el cuerpo del mensaje
	var bodyLines []string
	for scanner.Scan() {
		line := scanner.Text()
		bodyLines = append(bodyLines, line)
	}
	emailData.Body = strings.Join(bodyLines, "\n")

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &emailData, nil
}
