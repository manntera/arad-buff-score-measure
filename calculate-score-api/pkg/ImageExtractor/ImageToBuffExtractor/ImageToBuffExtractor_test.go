package Imagetobuffextractor

import (
	"context"
	"log"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
)

func TestGetSkillParams(t *testing.T) {
	ctx := context.Background()

	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	dir += "/../../../testdata/"

	for _, TestData := range Database.TestDataList {
		log.Printf("【Testing job】 %s", TestData.JobName)
		for _, imageData := range TestData.ImageDataList {
			log.Printf("【Testing image】 %s", imageData.ImageFileName)
			file, err := os.Open(dir + TestData.JobName + "/" + imageData.ImageFileName)
			if err != nil {
				t.Errorf("Error opening test image: %v", err)
				return
			}
			defer file.Close()

			text, err := ExtractBuffFromImage(ctx, file)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			t.Log(text)
		}
	}
}
