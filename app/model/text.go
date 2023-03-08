package model

type TextData struct {
	ID            string `json:"id"`
	Data          string `json:"textData"`
	ShouldEncrypt bool   `json:"encryption"`
	KeySize       uint   `json:"keySize"`
	PrivateKey    string `json:"privateKey"`
}
