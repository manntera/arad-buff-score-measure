package BuffSkillReader

import (
	"context"
	"log"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
	Imagetextextractor "manntera.com/calculate-score-api/pkg/ImageExtractor/ImageTextExtractor"
)

func TestGetBuffSkillParams(t *testing.T) {
	ctx := context.Background()
	rootDir, err := os.Getwd()
	rootDir += "/../../../testdata/"

	if err != nil {
		t.Errorf("Error getting working directory: %v", err)
		return
	}
	for _, TestData := range Database.TestDataList {
		log.Printf("【Testing job】 %s", TestData.JobName)
		for _, imageData := range TestData.ImageDataList {
			log.Printf("【Testing image】 %s", imageData.ImageFileName)
			imagePath := rootDir + TestData.JobName + "/" + imageData.ImageFileName
			file, err := os.Open(imagePath)
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
			if imageData.SkillId != buffSkillParams.SkillId {
				t.Errorf("Expected skill ID %d, got %d", imageData.SkillId, buffSkillParams.SkillId)
			}

			currentSuccessSkillParam := Database.GetSuccessParamFromSkillId(imageData.SkillId)
			if currentSuccessSkillParam == nil {
				t.Errorf("Expected success skill param, got nil")
			}
			for _, testParam := range buffSkillParams.BuffParams {
				SuccessParam := Database.GetSuccessParamFromparamId(*currentSuccessSkillParam, testParam.ParamId)
				if SuccessParam == nil {
					continue
				}
				if SuccessParam.ParamValue != testParam.ParamValue {
					t.Errorf("Expected param value %f, got %f", SuccessParam.ParamValue, testParam.ParamValue)
				}
			}
		}
	}
}
