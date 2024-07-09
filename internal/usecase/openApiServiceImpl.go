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

	const numWorkers = 5

	var wg sync.WaitGroup
	jobs := make(chan [2]string, len(hmapChapterContent))
	results := make(chan [2]string, len(hmapChapterContent))
	res = make(map[string]string, len(hmapChapterContent))

	// Khởi động các worker
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Gửi công việc đến các worker
	for key, v := range hmapChapterContent {
		jobs <- [2]string{key, v}
	}
	close(jobs) // Đóng channel jobs để các worker biết rằng không còn công việc nào

	// Chờ tất cả các worker hoàn thành
	wg.Wait()
	close(results)

	// Thu thập kết quả từ các worker
	for result := range results {
		key, summary := result[0], result[1]
		res[key] = summary
	}

	return res, err
}

func worker(jobs <-chan [2]string, results chan<- [2]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		key, content := job[0], job[1]
		shortContent, err := makeSummarizeRequest(content)
		if err == nil {
			results <- [2]string{key, *shortContent}
		}
	}
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
