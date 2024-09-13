package helper

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/generative-ai-go/genai"
)

func GetTextFromFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	text := string(data)
	// fmt.Println("1........", text)
	cleanedText := strings.TrimSpace(text)
	// fmt.Println("2........", cleanedText)

	return cleanedText, nil
}

func GetName(advice *genai.ChatSession, filePath string) (string, error) {
	text, err := GetTextFromFile(filePath)
	if err != nil {
		return "", err
	}

	request := fmt.Sprintf(`Based on the content of this document: %s, 
								provide only the full name 
									(first and last name) 
								of the author. 
							Do not include any additional information.`,
		text)

	resp, err := advice.SendMessage(context.Background(), genai.Text(request))
	if err != nil {
		return "", err
	}

	var output string
	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			output += fmt.Sprintf("%v", part)
		}
	}

	cleanedOutput := strings.TrimSpace(output)

	return cleanedOutput, nil
}