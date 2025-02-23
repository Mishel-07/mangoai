package mangoai

import "github.com/Mishel-07/mangoai/options"


func NewMango(ops ...options.Option) *options.Mango {
        mango := &options.Mango{
                baseURL: "https://mangooapi.onrender.com",
        }

        for _, option := range ops {
                option(mango)
        }

        mango.Chat = NewChat(mango)
        return mango
}
