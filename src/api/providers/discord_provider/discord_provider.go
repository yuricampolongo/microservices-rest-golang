package discord_provider

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/yuricampolongo/microservices-rest-golang/src/api/clients/restclient"
	"github.com/yuricampolongo/microservices-rest-golang/src/api/domain/discord"
)

const (
	// I'm leaving this hardcoded because if you want to fork or clone this repo, it'll be easier for you to test and see the results
	// so please, do not use this endpoint for spam or other purposes use only for studies and research.
	// If I begin to receive spams in my discord, I'll disable this endpoint and you have to create one on your own.
	// If you want to see the results you can join the discord server: https://discord.gg/kH8k5TDWd4 and check the channel #message-test
	urlSendMessage            = "https://discord.com/api/webhooks/827375998764449832/aQugmbNMF229HqYNKVcFMKIU6PqrJgkSJ3Zd17fs-46Z2nAJzT_wcWgnEjCdonkBwkYH"
	discordSucessResponseCode = 204
)

func SendMessage(message discord.DiscordMessage) (*discord.DiscordMessageResponse, *discord.DiscordErrorResponse) {
	response, err := restclient.Post(urlSendMessage, message)
	if err != nil {
		return nil, &discord.DiscordErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "error to send message to channel",
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &discord.DiscordErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "invalid body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode != discordSucessResponseCode {
		var errResponse discord.DiscordErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &discord.DiscordErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "invalid json error response body",
			}
		}
		return nil, &errResponse
	}

	var result discord.DiscordMessageResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, &discord.DiscordErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "invalid json response body",
		}
	}

	return &result, nil
}
