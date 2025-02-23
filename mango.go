package mangoai

import "github.com/Mishel-07/mangoai/options"

type Mango struct {
        baseURL string
        Chat    *Chat
}

func NewMango(options ...options.Option) *Mango {
        mango := &Mango{
                baseURL: "https://mangooapi.onrender.com",
        }

        for _, option := range options {
                option(mango)
        }

        mango.Chat = NewChat(mango)
        return mango
}
