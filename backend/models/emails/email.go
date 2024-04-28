package	emails


// Email model structure
type Email struct {	
	_id					 	string 
	MessageID               string 
	Date                    string 
	From                    string 
	To                      string 
	Subject                 string 
	MimeVersion             string 
	ContentType             string 
	ContentTransferEncoding string 
	XFrom                   string 
	XTo                     string 
	Xcc                     string 
	Xbcc                    string 
	XFolder                 string 
	XOrigin                 string 
	XFileName               string 
	Body                    string 
}		

type CreateEmail struct {
	MessageID               string `json:"Message-ID"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	To                      string `json:"To"`
	Subject                 string `json:"Subject"`
	MimeVersion             string `json:"Mime-Version"`
	ContentType             string `json:"Content-Type"`
	ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
	XFrom                   string `json:"X-From"`
	XTo                     string `json:"X-To"`
	Xcc                     string `json:"X-cc"`
	Xbcc                    string `json:"X-bcc"`
	XFolder                 string `json:"X-Folder"`
	XOrigin                 string `json:"X-Origin"`
	XFileName               string `json:"X-FileName"`
	Body                    string `json:"body"`
}


// func (cmd *CreateEmail) validate() error{


// }