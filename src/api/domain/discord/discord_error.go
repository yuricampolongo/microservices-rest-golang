package discord

type DiscordErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
