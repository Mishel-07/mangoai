package mangoai

import "errors"


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

func (c *Completions) Create(model string, messages []Messages) (*ChatCompletion, error) {
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

type Choice struct {
	Message Messages `json:"message"`
}

func NewChoice(json map[string]interface{}) Choice {
	return Choice{
		Message: NewMessages(json["message"].(map[string]interface{})),
	}
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func NewMessages(json map[string]interface{}) Messages {
	return Messages{
		Role:    json["role"].(string),
		Content: json["content"].(string),
        }
}
