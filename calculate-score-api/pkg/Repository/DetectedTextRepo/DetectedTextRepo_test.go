package DetectedTextRepo

import (
	"context"
	"log"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/Repository/SamplerImageRepo"
)

func TestDetectedTextRepo(t *testing.T) {
	ctx := context.Background()
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	dir += "/../../../testdata/"
	for _, testData := range Database.TestDataList {
		log.Printf("【Testing job】 %s", testData.JobName)
		for _, imageData := range testData.ImageDataList {
			log.Printf("【Testing image】 %s", imageData.ImageFileName)
			testImagePath := dir + testData.JobName + "/" + imageData.ImageFileName
			file, err := os.Open(testImagePath)
			if err != nil {
				t.Errorf("Error opening test image: %v", err)
				continue
			}
			defer file.Close()
			samplerImageRepo, err := SamplerImageRepo.NewSamplerImageRepoFromFile(file)
			if err != nil {
				t.Errorf("Error creating SamplerImageRepo: %v", err)
				continue
			}
			defer samplerImageRepo.Close()

			detectedTextRepo, err := NewDetectedTextRepoFromSamplerImageRepo(ctx, samplerImageRepo)
			if err != nil {
				t.Errorf("Error creating DetectedTextRepo: %v", err)
				continue
			}
			log.Printf("Detected text: %v", detectedTextRepo)
			//detectedTextRepo.FindLineTextFromKeyword
		}
	}
}
