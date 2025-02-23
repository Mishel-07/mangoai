package mangoai

import "github.com/Mishel-07/mangoai/options"


func NewMango(options ...options.Option) *options.Mango {
        mango := &Mango{
                baseURL: "https://mangooapi.onrender.com",
        }

        for _, option := range options {
                option(mango)
        }

        mango.Chat = NewChat(mango)
        return mango
}
