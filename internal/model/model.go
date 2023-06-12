package model

type InputModel struct {
	Text string `json:"text"`
	Time string `json:"time"`
	Lang string `json:"language"` // 0 - English, 1 - Russian
	// ClientType string `json:"client_type"` // "web" or "desktop"
}

type OutputModel struct {
	Text  string `json:"text"`  //Response text
	Error string `json:"error"` //Error text
}

const (
	English = iota
	Russian
)
