package discord

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscordMessageRequestAsJson(t *testing.T) {
	request := Message{
		Content: "test message golang api",
		Tts:     false,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target Message

	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Content, request.Content)
}
