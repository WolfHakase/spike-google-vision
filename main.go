package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

const jsonPath = "credentials.json"
var filenames = []string {"testdata/ST1-11.png", "testdata/BT1-025.jpg"}

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	visionClient := Vision{client: client}

	for _, filename := range filenames {
		visionClient.PrintCardIDForFile(ctx, filename)
	}
}

type Vision struct {
	client *vision.ImageAnnotatorClient
}

func (v Vision) PrintCardIDForFile(ctx context.Context, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	texts, err := v.client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect texts: %v", err)
	}

	for _, text := range texts {
		if isStarterSet(text.Description) || isBoosterSet(text.Description) {
			fmt.Println("Card ID: ", text.Description)
		}
	}
}

func isStarterSet(text string) bool {
	return strings.Contains(text, "ST")
}

func isBoosterSet(text string) bool {
	return strings.Contains(text, "BT")
}
