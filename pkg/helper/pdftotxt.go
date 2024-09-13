package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pdfcrowd/pdfcrowd-go"
	"golang.org/x/exp/slog"
)

func readFile(fileName string) ([]byte, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		slog.Error("Error reading file: %v", err)
		return nil, err
	}
	return content, nil
}

func ConvertationToTXT(path string) (string, error) {
	client := pdfcrowd.NewPdfToTextClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")

	outputFilePath := filepath.Join("internal", "media", "invoice.txt")
	outputStream, err := os.Create(outputFilePath)
	if err != nil {
		return "", fmt.Errorf("error creating output file: %v", err)
	}
	defer outputStream.Close()

	fileContent, err := readFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading PDF file: %v", err)
	}

	err = client.ConvertRawDataToStream(fileContent, outputStream)
	if err != nil {
		return "", fmt.Errorf("error converting PDF to text: %v", err)
	}

	return outputFilePath, nil
}
