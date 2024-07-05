package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"scaleX/internal/constants"
	"scaleX/internal/dto"
	"strings"
	"sync"
)

type openApiService struct {
}

func (b openApiService) SummarizeChapters(ctx context.Context) (res map[string]string, err error) {
	fileContent, err := os.ReadFile("./sampleFile/TheArtOfThinkingClearly.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	content := string(fileContent)
	hmapChapterContent := getChapterAndContent(content)

	var wg sync.WaitGroup
	wg.Add(len(hmapChapterContent))

	res = make(map[string]string, len(hmapChapterContent))
	for key, v := range hmapChapterContent {
		k := key
		v := v
		go func() {
			defer wg.Done()
			shortContent, err := makeSummarizeRequest(v)
			if err == nil {
				res[k] = *shortContent
			}
		}()
	}

	wg.Wait()

	return res, err
}

func makeSummarizeRequest(chapterContent string) (shortContent *string, err error) {
	apiKey := "sk-team2024-home-test-vP00R3HgNtYAroFA1rRbT3BlbkFJ6XTugV2XnCDwPKOnwg18"

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
		Content: chapterContent,
	}

	reqBody.Messages = append(reqBody.Messages, systemMsg, userMsg)

	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Failed to marshal request body:", err)
		return nil, err
	}

	// Tạo HTTP request
	req, err := http.NewRequest("POST", constants.OpenaiAPIURL, bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Gửi request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Đọc và phân tích response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return nil, err
	}

	var content dto.OpenAIResponse
	err = json.Unmarshal(respBody, &content)
	if err != nil {
		fmt.Println("Failed to unmarshal response body:", err)
		return nil, err
	}

	return &content.Choices[0].Message.Content, err
}

func getChapterAndContent(allContent string) (res map[string]string) {
	re := regexp.MustCompile(`(?m)^(\d+)\s*\n`)

	splitContent := re.Split(allContent, -1)

	headers := re.FindAllStringSubmatch(allContent, -1)

	chapters := make(map[string]string)

	for i, header := range headers {
		chapters[header[1]] = strings.TrimSpace(splitContent[i+1])
	}

	res = make(map[string]string, len(chapters))

	for chapter, text := range chapters {
		res["Chapter "+chapter] = text
	}

	return res
}

func NewOpenApiService() OpenApiService {
	return &openApiService{}
}
