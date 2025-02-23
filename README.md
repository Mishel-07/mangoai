# Mango
Mango 🥭 is a Golang wrapper for the MangoAPI, providing powerful AI functionalities and LLM capabilities. 
It's really cool! Give it a try and access some of its models for free.


## Installation

Install the mango api
```bash
go get github.com/Mishel-07/mangoai@latest
```

## Usage

Here is a basic example of how to use mango ai:

```python
package main

import (
 "fmt"
 "github.com/Mishel-07/mangoai/mango"
)

func main() {
 client := mango.NewMango()

 messages := []mango.Messages{
  {
   Role:    "user",
   Content: "hi",
  },
 }

 chatCompletion, err := client.Chat.Completions.Create("gpt-3.5-turbo", messages)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }

 if len(chatCompletion.Choices) > 0 {
  content := chatCompletion.Choices[0].Message.Content
  fmt.Println("Response Content:", content)
 } else {
  fmt.Println("No choices available")
 }
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

## Contact
For support or inquiries, please reach out to us at [Telegram Support](https://t.me/XBOTSUPPORTS).
