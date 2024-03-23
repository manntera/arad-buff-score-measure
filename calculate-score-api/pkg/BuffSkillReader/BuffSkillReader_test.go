package BuffSkillReader

import (
	"context"
	"os"
	"testing"

	Imagetextextractor "manntera.com/calculate-score-api/pkg/ImageTextExtractor"
)

func TestGetBuffSkillParams(t *testing.T) {
	ctx := context.Background()
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	testImagePath := dir + "/../../testdata/test2.jpg"

	file, err := os.Open(testImagePath)
	if err != nil {
		t.Errorf("Error opening test image: %v", err)
		return
	}
	defer file.Close()

	text, err := Imagetextextractor.ExtractTextFromImage(ctx, file)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	buffSkillParams, buffSkillParamsErr := GetBuffSkillParams(text)
	if buffSkillParamsErr != nil {
		t.Errorf("Error: %v", buffSkillParamsErr)
	}
	if len(buffSkillParams) == 0 {
		t.Errorf("Expected buffSkillParams to be non-empty")
	}
	t.Log(buffSkillParams)
}
