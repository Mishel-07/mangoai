package mangoai

type Mango struct {
        baseURL string
        Chat    *Chat
}

type Option func(*Mango)

func WithBaseURL(url string) Option {
        return func(m *Mango) {
                m.baseURL = url
        }
}

func NewMango(options ...Option) *Mango {
        mango := &Mango{
                baseURL: "https://www.api.mangoi.in",
        }

        for _, option := range options{
                option(mango)
        }

        mango.Chat = NewChat(mango)
        return mango
}
