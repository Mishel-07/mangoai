package options

import "github.com/Mishel-07/mangoai"

type Option func(*mangoai.Mango)

const RoleUser = "user"

const RoleSystem = "system"

const Assistant = "assistant"

func WithBaseURL(url string) Option {
        return func(m *mangoai.Mango) {
            m.baseURL = url
        }
}
