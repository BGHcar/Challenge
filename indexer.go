package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Email struct {
	MessageID           string                 `json:"Message-ID"`
	Date                string                 `json:"Date"`
	From                string                 `json:"From"`
	To                  string                 `json:"To"`
	Subject             string                 `json:"Subject"`
	MimeVersion         string                 `json:"Mime-Version"`
	ContentType         string                 `json:"Content-Type"`
	ContentTransfer     string                 `json:"Content-Transfer-Encoding"`
	Content             string                 `json:"Content"`
	AdditionalProperties map[string]interface{} `json:"additionalProperties"`
}

func main() {
	// Cargar las variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Obtener la ruta del directorio base desde las variables de entorno
	baseDir := os.Getenv("PATH_DIRECTION")
	if baseDir == "" {
		log.Fatal("La variable de entorno PATH_DIRECTION no está definida")
	}

	// Directorio donde se guardará el archivo JSON
	outputPath := "output.json"

	// Indexar archivos
	emails := make(chan Email)
	var wg sync.WaitGroup
	var count int
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == "" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				email, err := parseEmailFile(path)
				if err != nil {
					log.Printf("Error parsing file %s: %v", path, err)
					return
				}
				emails <- email
			}()
			log.Printf("total de correos procesados: %d", count)
			count++
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking the path %s: %v", baseDir, err)
	}

	go func() {
		wg.Wait()
		close(emails)
	}()

	// Crear la estructura de datos con el formato deseado
	data := struct {
		Index   string  `json:"index"`
		Records []Email `json:"records"`
	}{
		Index:   "email",
		Records: []Email{},
	}

	for email := range emails {
		data.Records = append(data.Records, email)
	}

	// Escribir los datos en un archivo JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling data to JSON: %v", err)
	}

	err = ioutil.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing data to file: %v", err)
	}

	fmt.Printf("Se han procesado y guardado %d registros en %s\n", count, outputPath)
}

func parseEmailFile(path string) (Email, error) {
	file, err := os.Open(path)
	if err != nil {
		return Email{}, err
	}
	defer file.Close()

	// Crear un nuevo escáner con un búfer más grande
	scanner := bufio.NewScanner(file)
	const maxCapacity = 1024 * 1024 // Tamaño máximo del búfer en bytes
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	email := Email{}
	var content strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Message-ID: ") {
			email.MessageID = strings.TrimPrefix(line, "Message-ID: ")
		} else if strings.HasPrefix(line, "Date: ") {
			email.Date = strings.TrimPrefix(line, "Date: ")
		} else if strings.HasPrefix(line, "From: ") {
			email.From = strings.TrimPrefix(line, "From: ")
		} else if strings.HasPrefix(line, "To: ") {
			email.To = strings.TrimPrefix(line, "To: ")
		} else if strings.HasPrefix(line, "Subject: ") {
			email.Subject = strings.TrimPrefix(line, "Subject: ")
		} else if strings.HasPrefix(line, "Mime-Version: ") {
			email.MimeVersion = strings.TrimPrefix(line, "Mime-Version: ")
		} else if strings.HasPrefix(line, "Content-Type: ") {
			email.ContentType = strings.TrimPrefix(line, "Content-Type: ")
		} else if strings.HasPrefix(line, "Content-Transfer-Encoding: ") {
			email.ContentTransfer = strings.TrimPrefix(line, "Content-Transfer-Encoding: ")
		} else {
			// Asumimos que el contenido del correo comienza después de las cabeceras
			content.WriteString(line)
			content.WriteString("\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return Email{}, err
	}

	email.Content = content.String()

	// Agregar propiedades adicionales al correo electrónico, como la fecha de creación
	email.AdditionalProperties = map[string]interface{}{
		"@timestamp": time.Now().Format(time.RFC3339),
		"_id":        filepath.Base(path),
		"x_filename": filepath.Base(path), // A modo de ejemplo, guardamos el nombre del archivo como propiedad adicional
	}

	return email, nil
}
