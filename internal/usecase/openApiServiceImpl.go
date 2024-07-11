package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
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
	jobs := make(chan dto.ChapterContent, len(hmapChapterContent))
	results := make(chan dto.ChapterContent, len(hmapChapterContent))
	res = make(map[string]string, len(hmapChapterContent))

	// Create 5 worker
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// Push work into workers
	for key, v := range hmapChapterContent {
		chapterContent := dto.ChapterContent{
			Chapter: key,
			Content: v,
		}
		jobs <- chapterContent
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		key, summary := result.Chapter, result.Content
		res[key] = summary
	}

	return res, err
}

func worker(jobs <-chan dto.ChapterContent, results chan<- dto.ChapterContent, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		log.Info("Doing job with chapter ", job.Chapter)
		shortContent, err := makeSummarizeRequest(job.Content)
		job.Content = *shortContent
		if err == nil {
			results <- job
			log.Info("Done job with chapter ", job.Chapter)
		} else {
			log.Errorf("makeSummarizeRequest.err %v", err)
		}
	}
}

func makeSummarizeRequest(chapterContent string) (shortContent *string, err error) {
	apiKey := "sk-team2024-home-test-vP00R3HgNtYAroFA1rRbT3BlbkFJ6XTugV2XnCDwPKOnwg18" // need to put it on .env

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
		log.Errorf("Failed to marshal request body: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", constants.OpenaiAPIURL, bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		log.Error("Failed to create request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed to read response body:", err)
		return nil, err
	}

	var content dto.OpenAIResponse
	err = json.Unmarshal(respBody, &content)
	if err != nil {
		log.Error("Failed to unmarshal response body:", err)
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
