package Imagetextextractor

import (
	"context"
	"log"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
)

func TestExtractTextFromImage(t *testing.T) {
	ctx := context.Background()
	dir, err := os.Getwd()
	dir += "/../../../testdata/"
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	for _, testData := range Database.TestDataList {
		log.Printf("【Testing job】 %s", testData.JobName)
		for _, imageData := range testData.ImageDataList {
			log.Printf("【Testing image】 %s", imageData.ImageFileName)
			testImagePath := dir + testData.JobName + "/" + imageData.ImageFileName
			file, err := os.Open(testImagePath)
			if err != nil {
				t.Errorf("Error opening test image: %v", err)
				return
			}
			defer file.Close()
			text, err := ExtractTextFromImage(ctx, file)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if text == "" {
				t.Errorf("Expected text to be non-empty")
			}
			log.Printf("【Extracted text】 %s", text)
		}
	}
}
