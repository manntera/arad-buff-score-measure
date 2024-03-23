package CalculateBuffScoreUsecase

import (
	"context"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Usecase/AnalyzeBuffSkillsFromImageUsecase"
)

func TestCalculateBuffScore(t *testing.T) {
	ctx := context.Background()
	fileNames := []string{
		"test1.jpg",
		"test2.jpg",
		"test3.jpg",
		"test4.jpg",
		"test5.png",
	}
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	testDir := dir + "/../../../testdata/"
	var files []os.File
	for _, fileName := range fileNames {
		file, err := os.Open(testDir + fileName)
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

	if score < 0 {
		t.Errorf("Expected score to be non-negative")
	}
	t.Log("BuffScore", score)
}
