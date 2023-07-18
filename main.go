// REF https://qiita.com/bluewolfali/items/9031cd56f1743aec3e9a

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

var messages []Message

func main() {
	// API_KEYを設定してください
	apiKey := "API_KEY"

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("ワイルドなご意見を入力してください。マイルドにします。: ")
		opinion, _ := reader.ReadString('\n')
		opinion = strings.TrimSpace(opinion)

		if opinion == "exit" {
			break
		}

		opinion += "「" + opinion + "」の言い方をマイルドにしてください。"

		messages = append(messages, Message{
			Role:    "user",
			Content: opinion,
		})

		response := getOpenAIResponse(apiKey)
		fmt.Println("【マイルド】" + response.Choices[0].Messages.Content)
		print("\n")
	}
}

func getOpenAIResponse(apiKey string) OpenaiResponse {
	requestBody := OpenaiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	requestJSON, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response OpenaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		println("Error: ", err.Error())
		return OpenaiResponse{}
	}

	messages = append(messages, Message{
		Role:    "assistant",
		Content: response.Choices[0].Messages.Content,
	})

	return response
}
