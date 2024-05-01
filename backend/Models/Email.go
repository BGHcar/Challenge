package Models

type Email struct {
	Id       string   `json:"_id"`
	From     string   `json:"From"`
	To       string   `json:"To"`
	Subject  string   `json:"Subject"`
	Metadata Metadata `json:"Metadata"`
	Message  string   `json:"Message"`
} // Definición de la estructura Email

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
