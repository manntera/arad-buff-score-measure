package BuffParamReader

import (
	"context"
	"log"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
	Imagetextextractor "manntera.com/calculate-score-api/pkg/ImageExtractor/ImageTextExtractor"
)

func TestGetBuffParam(t *testing.T) {
	ctx := context.Background()
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	for _, TestData := range Database.TestDataList {
		log.Printf("【Testing job】 %s", TestData.JobName)
		testImagePath := dir + "/../../../testdata/" + TestData.JobName
		for _, imageData := range TestData.ImageDataList {
			log.Printf("【Testing image】 %s", imageData.ImageFileName)
			testImageFilePath := testImagePath + "/" + imageData.ImageFileName
			file, err := os.Open(testImageFilePath)
			if err != nil {
				t.Errorf("Error opening test image: %v", err)
				return
			}
			defer file.Close()
			text, err := Imagetextextractor.ExtractTextFromImage(ctx, file)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			currentSkill := Database.GetSkillFromId(imageData.SkillId)

			if currentSkill == nil {
				t.Errorf("Expected skill, got nil")
				continue
			}

			buffParam, buffParamErr := GetBuffParams(text, currentSkill.Name)
			if buffParamErr != nil {
				t.Errorf("Error: %v", buffParamErr)
			}
			if len(buffParam) == 0 {
				t.Errorf("Expected buffParam to be non-empty")
			}

			for _, testParam := range buffParam {
				currentSuccessSkillParam := Database.GetSuccessParamFromSkillId(imageData.SkillId)
				if currentSuccessSkillParam == nil {
					t.Errorf("Expected success skill param, got nil")
					continue
				}
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
