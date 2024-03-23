package Imagetobuffextractor

import (
	"context"
	"os"
	"testing"
)

func TestGetSkillParams(t *testing.T) {
	ctx := context.Background()
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	// testImagePath := dir + "/../../testdata/test5.png"
	testImagePath := dir + "/../../testdata/test1.jpg"
	file, err := os.Open(testImagePath)
	if err != nil {
		t.Errorf("Error opening test image: %v", err)
		return
	}
	defer file.Close()

	buffSkillParams, buffSkillParamsErr := ExtractBuffFromImage(ctx, file)
	if buffSkillParamsErr != nil {
		t.Errorf("Error: %v", buffSkillParamsErr)
	}
	if len(buffSkillParams) == 0 {
		t.Errorf("Expected buffSkillParams to be non-empty")
	}
	t.Log(buffSkillParams)
}
