package discord

type Message struct {
	Content string `json:"content"`
	Tts     bool   `json:"tts"`
}

type MessageResponse struct {
	Code int `json:"code"`
}
