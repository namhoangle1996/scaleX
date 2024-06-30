package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"scaleX/internal/constants"
	"scaleX/internal/dto"
)

// URL API OpenAI

func main() {
	apiKey := "sk-team2024-home-test-vP00R3HgNtYAroFA1rRbT3BlbkFJ6XTugV2XnCDwPKOnwg18"

	// Văn bản chương mẫu để tóm tắt
	chapterText := `
1
WHY YOU SHOULD VISIT CEMETERIES
Survivorship Bias
No matter where Rick looks, he sees rock stars. They appear on television, on the front pages of magazines, in concert programmes and at online fan sites. Their songs are unavoidable - in the mall, on his playlist, in the gym. The rock stars are everywhere. There are lots of them. And they are successful. Motivated by the stories of countless guitar heroes, Rick starts a band. Will he make it big? The probability lies a fraction above zero. Like so many others, he will most likely end up in the graveyard of failed musicians. This burial ground houses 10,000 times more musicians than the stage does, but no journalist is interested in failures - with the exception of fallen superstars. This makes the cemetery invisible to outsiders.
In daily life, because triumph is made more visible than failure, you systematically overestimate your chances of succeeding. As an outsider, you (like Rick) succumb to an illusion, and you mistake how minuscule the probability of success really is. Rick, like so many others, is a victim of Survivorship Bias.
Behind every popular author you can find 100 other writers whose books will never sell. Behind them are another 100 who haven't found publishers. Behind them are yet another 100 whose unfinished manuscripts gather dust in drawers.
And behind each one of these are 100 people who dream of - one day - writing a book. You, however, hear of only the successful authors (these days, many of them self-published) and fail to recognise how unlikely literary success is. The same goes for photographers, entrepreneurs, artists, athletes, architects, Nobel Prize winners, television presenters and beauty queens. The media is not interested in digging around in the graveyards of the unsuccessful. Nor is this its job. To elude the survivorship bias, you must do the digging yourself.
You will also come across survivorship bias when dealing with money and risk: imagine that a friend founds a start-up. You belong to the circle of potential investors and you sense a real opportunity: this could be the next Google. Maybe you'll be lucky. But what is the reality? The most likely scenario is that the company will not even make it off the starting line.
`
	// Tạo payload request
	reqBody := dto.OpenAIRequest{
		Model:       "gpt-3.5-turbo",
		MaxTokens:   88,
		Temperature: 0.7,
	}

	systemMsg := dto.OpenAIMessage{
		Role:    "system",
		Content: "Summarize content you are provided with for a second-grade student.",
	}

	userMsg := dto.OpenAIMessage{
		Role:    "user",
		Content: chapterText,
	}

	reqBody.Messages = append(reqBody.Messages, systemMsg, userMsg)

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Failed to marshal request body:", err)
		return
	}

	fmt.Println("reqBodyJSON", string(reqBodyJSON))

	// Tạo HTTP request
	req, err := http.NewRequest("POST", constants.OpenaiAPIURL, bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Gửi request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	// Đọc và phân tích response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	var content dto.OpenAIResponse
	err = json.Unmarshal(respBody, &content)
	if err != nil {
		panic(err)
	}
	fmt.Println("resp.Content", content.Choices[0].Message.Content)
}
