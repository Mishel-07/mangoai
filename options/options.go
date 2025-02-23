package options

type Mango struct {
        baseURL string
        Chat    interface{}
}

type Option func(*Mango)

const RoleUser = "user"

const RoleSystem = "system"

const RoleAssistant = "assistant"

func WithBaseURL(url string) Option {
        return func(m *mangoai.Mango) {
            m.baseURL = url
        }
}
