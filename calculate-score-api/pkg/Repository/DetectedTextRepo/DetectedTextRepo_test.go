package DetectedTextRepo

import (
	"context"
	"encoding/json"
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
			testDir := dir + testData.JobName + "/"
			testImagePath := testDir + imageData.ImageFileName
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
			json, err := json.MarshalIndent(detectedTextRepo, "", "  ")
			if err != nil {
				t.Errorf("Error marshalling DetectedTextRepo: %v", err)
				continue
			}
			jsonFilePath := testDir + "detectedText/" + imageData.ImageFileName + ".json"
			os.Mkdir(testDir+"detectedText/", 0777)
			os.Create(jsonFilePath)

			err = os.WriteFile(jsonFilePath, json, 0666)
			if err != nil {
				t.Errorf("Error writing json file: %v", err)
				continue
			}
		}
	}
}
