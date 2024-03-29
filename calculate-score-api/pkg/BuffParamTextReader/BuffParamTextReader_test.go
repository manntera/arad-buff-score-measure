package BuffParamTextReader

import (
	"context"
	"log"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
	ImageTextExtractor "manntera.com/calculate-score-api/pkg/ImageExtractor/ImageTextExtractor"
)

func TestGetBuffParam(t *testing.T) {
	ctx := context.Background()
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	for _, TestData := range Database.TestDataList {
		currentJobName := TestData.JobName
		log.Printf("【Testing job】 %s", currentJobName)
		testImagePath := dir + "/../../testdata/" + currentJobName
		for _, testImageData := range TestData.ImageDataList {
			log.Printf("【Testing image】 %s", testImageData.ImageFileName)
			testImageFilePath := testImagePath + "/" + testImageData.ImageFileName
			file, err := os.Open(testImageFilePath)
			if err != nil {
				t.Errorf("Error opening test image: %v", err)
				return
			}
			defer file.Close()
			currentText, err := ImageTextExtractor.ExtractTextFromImage(ctx, file)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			currentBuffParam, buffParamErr := GetBuffParam(currentText)
			if buffParamErr != nil {
				t.Errorf("Error: %v", buffParamErr)
				return
			}
			if currentBuffParam.SkillId != testImageData.SkillId {
				t.Errorf("Expected skill id %d, got %d", testImageData.SkillId, currentBuffParam.SkillId)
			}
			if testImageData.BaseParam != currentBuffParam.BaseParam {
				t.Errorf("Expected base param %d, got %d", testImageData.BaseParam, currentBuffParam.BaseParam)
			}
			if testImageData.BoostParam != currentBuffParam.BoostParam {
				t.Errorf("Expected boost param %d, got %d", testImageData.BoostParam, currentBuffParam.BoostParam)
			}
		}
	}
}
