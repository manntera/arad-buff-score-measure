package CalculateBuffScore

import (
	"context"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/Usecase/AnalyzeBuffSkillsFromImageUsecase"
)

func TestCalculateBuffScore(t *testing.T) {
	ctx := context.Background()

	rootDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting working directory: %v", err)
		return
	}
	rootDir += "/../../testdata/"
	for _, testData := range Database.TestDataList {
		imageDir := rootDir + testData.JobName + "/"
		var files []os.File
		for _, imageData := range testData.ImageDataList {
			file, err := os.Open(imageDir + imageData.ImageFileName)
			if err != nil {
				t.Errorf("Error opening test image: %v", err)
				return
			}
			defer file.Close()
			files = append(files, *file)
		}
		buffSkillParams, err := AnalyzeBuffSkillsFromImageUsecase.AnalyzeBuffSkillsFromImages(ctx, files)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		score, err := CalculateBuffScore(buffSkillParams)
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if score != testData.Score {
			t.Errorf("Expected score %v, got %v", testData.Score, score)
		}
		if score == testData.Score {
			t.Logf("Score is correct")
		}
	}
}
