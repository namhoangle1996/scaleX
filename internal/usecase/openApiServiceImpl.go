package usecase

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type openApiService struct {
}

func (b openApiService) SummarizeChapters(ctx context.Context) error {
	fileContent, err := os.ReadFile("./sampleFile/TheArtOfThinkingClearly.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	content := string(fileContent)

	re := regexp.MustCompile(`(?m)^(\d+)\s*\n`)

	splitContent := re.Split(content, -1)
	fmt.Println("splitContent", splitContent)

	headers := re.FindAllStringSubmatch(content, -1)

	// Create a map to store chapters
	chapters := make(map[string]string)

	// Iterate over the split content and headers
	for i, header := range headers {
		chapters[header[1]] = strings.TrimSpace(splitContent[i+1])
	}

	for chapter, text := range chapters {
		fmt.Printf("Chapter %s:\n%s\n\n", chapter, text)
	}

}

func NewOpenApiService() OpenApiService {
	return &openApiService{}
}
