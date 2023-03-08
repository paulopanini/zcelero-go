package model

type TextData struct {
	ID         int    `json:"id"`
	Data       string `json:"text_data"`
	Encryption bool   `json:"encryption"`
	KeySize    uint   `json:"key_size"`
}
