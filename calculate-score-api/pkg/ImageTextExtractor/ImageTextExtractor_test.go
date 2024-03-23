package Imagetextextractor

import (
	"context"
	"os"
	"testing"
)

func TestExtractTextFromImage(t *testing.T) {
	ctx := context.Background()
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	testImagePath := dir + "/../../../testdata/test4.jpg"

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
	t.Log(text)
}
