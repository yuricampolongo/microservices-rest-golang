package discord

type DiscordMessage struct {
	Content string       `json:"content"`
	Tts     bool         `json:"tts"`
	Embed   EmbedMessage `json:"embed"`
}

type EmbedMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DiscordMessageResponse struct {
	Code int `json:"code"`
}
