package ImageProcessing

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"

	"manntera.com/calculate-score-api/pkg/BuffParamTextReader"
	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/ImageExtractor/ImageTextExtractor"
)

func TestTrimBuffIconArea(t *testing.T) {
	ctx := context.Background()
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}

	for _, TestData := range Database.TestDataList {
		log.Printf("====================")
		currentJobName := TestData.JobName
		log.Printf("【Testing job】 %s", currentJobName)
		testImagePath := dir + "/../../testdata/" + currentJobName
		for _, testImageData := range TestData.ImageDataList {
			log.Printf("-------")
			log.Printf("【Testing image】 %s", testImageData.ImageFileName)
			testImageFilePath := testImagePath + "/" + testImageData.ImageFileName
			file, err := os.Open(testImageFilePath)
			if err != nil {
				t.Errorf("Error opening test image: %v", err)
				return
			}
			defer file.Close()

			extension := filepath.Ext(testImageData.ImageFileName)
			// trimedImage := file
			trimedImage := new(os.File)
			switch extension {
			case ".png":
				trimedImage, err = TrimBuffIconArea(PNG, file)
			case ".jpg":
				trimedImage, err = TrimBuffIconArea(JPG, file)
			default:
				t.Errorf("Invalid image type: %s", extension)
			}
			if err != nil {
				t.Errorf("Error trimming image: %v", err)
				return
			}
			currentText, err := ImageTextExtractor.ExtractTextFromImage(ctx, trimedImage)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			currentBuffParam, buffParamErr := BuffParamTextReader.GetBuffParam(currentText)
			if buffParamErr != nil {
				t.Logf("【Extracted text】\n%s", currentText)
				t.Errorf("Error: %v", buffParamErr)
				continue
			}
			if currentBuffParam.SkillId != testImageData.SkillId {
				t.Errorf("Expected skill id %d, got %d", testImageData.SkillId, currentBuffParam.SkillId)
			}
			if testImageData.BaseParam != currentBuffParam.BaseParam {
				t.Errorf("Expected base param %d, got %d", testImageData.BaseParam, currentBuffParam.BaseParam)
			}
			if testImageData.BoostParam != currentBuffParam.BoostParam {
				t.Logf("【Extracted text】\n%s", currentText)
				t.Errorf("Expected boost param %d, got %d", testImageData.BoostParam, currentBuffParam.BoostParam)
			}
			skill, err := Database.GetSkillFromId(testImageData.SkillId)

			name := ""

			if err != nil {
				name = "error"
			} else {
				name = skill.Name
			}

			newFileName := "trimmed_" + name + ".png"
			newFilePath := filepath.Join(filepath.Dir(trimedImage.Name()), newFileName)
			err = os.Rename(trimedImage.Name(), newFilePath)

			if err != nil {
				t.Errorf("Error renaming file: %v", err)
			}
			trimedImage.Close()
		}
	}
}
