package Models

type Email struct {
	Id     string `json:"-"`
	From    string `json:"From"`
	To      string `json:"To"`
	Subject string `json:"Subject"`
	Body    string `json:"Body"`
} // Definición de la estructura Email
