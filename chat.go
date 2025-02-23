package mangoai

import (
	"errors"
)

type Chat struct {
	mango       *Mango
	Completions *Completions
}

func NewChat(mango *Mango) *Chat {
	chat := &Chat{mango: mango}
	chat.Completions = &Completions{chat: chat}
	return chat
}

type Completions struct {
	chat *Chat
}

func (c *Completions) Create(model string, messages []Message) (*ChatCompletion, error) {
	if model == "" {
		return nil, errors.New("model is required, You can see model here https://mangooapi.onrender.com/models")
	}
	if len(messages) == 0 {
		return nil, errors.New("messages is required")
	}

	payload := map[string]interface{}{
		"model":    model,
		"messages": messages,
	}

	response, err := c.chat.mango.DoRequest("mango", "POST", payload)
	if err != nil {
		return nil, err
	}

	if response["status"] == "error" {
		return nil, errors.New("Error: Report https://github.com/Mishel-07/mangoai/issues")
	}
	if response["response"] == "invalid model" {
		return nil, errors.New("Invalid model")
	}

	return NewChatCompletion(response), nil
}

type ChatCompletion struct {
	Status   string    `json:"status"`
	Object   string    `json:"object"`
	Response string    `json:"response"`
	Choices  []Choice  `json:"choices"`
}

func NewChatCompletion(response map[string]interface{}) *ChatCompletion {
	completion := &ChatCompletion{
		Status:   response["status"].(string),
		Object:   response["object"].(string),
		Response: response["response"].(string),
	}

	var choices []Choice
	for _, choice := range response["choices"].([]interface{}) {
		choices = append(choices, NewChoice(choice.(map[string]interface{})))
	}
	completion.Choices = choices

	return completion
}

type Choice struct {
	Message Message `json:"message"`
}

func NewChoice(json map[string]interface{}) Choice {
	return Choice{
		Message: NewMessage(json["message"].(map[string]interface{})),
	}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewMessage(json map[string]interface{}) Message {
	return Message{
		Role:    json["role"].(string),
		Content: json["content"].(string),
	}
}
