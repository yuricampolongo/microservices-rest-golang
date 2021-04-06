package discord

type Message struct {
	Content string       `json:"content"`
	Tts     bool         `json:"tts"`
	Embed   EmbedMessage `json:"embed"`
}

type EmbedMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MessageResponse struct {
	Code int `json:"code"`
}
