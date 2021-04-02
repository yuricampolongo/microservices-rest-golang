package discord

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscordMessageRequestAsJson(t *testing.T) {
	request := DiscordMessage{
		Content: "test message golang api",
		Tts:     false,
		Embed: EmbedMessage{
			Title:       "Test message golang embed",
			Description: "this is an embed message",
		},
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target DiscordMessage

	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Content, request.Content)
}
