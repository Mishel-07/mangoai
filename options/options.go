package options

import "github.com/Mishel-07/mangoai"

type Option func(*mangoai.Mango)

func WithBaseURL(url string) Option {
        return func(m *mangoai.Mango) {
            m.baseURL = url
        }
}
