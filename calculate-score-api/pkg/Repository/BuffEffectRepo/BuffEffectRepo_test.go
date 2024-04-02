package BuffEffectRepo

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/Repository/BuffSkillRepo"
	"manntera.com/calculate-score-api/pkg/Repository/DetectedTextRepo"
)

func TestBuffEffectRepo(t *testing.T) {
	baseDir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	testdataDir := baseDir + "/../../../testdata/"
	skillFileDir := baseDir + "/../../../setting/BuffSkill.json"
	skillRepo, err := BuffSkillRepo.NewBuffSkillRepoFromJsonFile(skillFileDir)
	if err != nil {
		t.Errorf("Failed to create a new BuffSkillRepo: %v", err)
	}
	for _, testData := range Database.TestDataList {
		log.Printf("【Testing job】 %s", testData.JobName)
		for _, imageData := range testData.ImageDataList {
			log.Printf("【Testing image】 %s", imageData.ImageFileName)
			testDir := testdataDir + testData.JobName + "/"
			testJsonFile := testDir + "detectedText/" + imageData.ImageFileName + ".json"
			bytes, err := os.ReadFile(testJsonFile)
			if err != nil {
				t.Errorf("Error reading json file: %v", err)
				continue
			}
			var detectedTextRepo *DetectedTextRepo.DetectedTextRepo
			err = json.Unmarshal(bytes, &detectedTextRepo)
			if err != nil {
				t.Errorf("Error unmarshalling json file: %v", err)
				continue
			}
			buffEffectRepo, err := NewBuffEffectRepoFromDetectedTextRepo(skillRepo, detectedTextRepo)
			if err != nil {
				t.Errorf("Error creating BuffEffectRepo: %v", err)
				continue
			}
			if buffEffectRepo.buffEffects.BaseParam != imageData.BaseParam {
				t.Errorf("BaseParam not match: %v", buffEffectRepo.buffEffects.BaseParam)
			}
			if buffEffectRepo.buffEffects.BoostParam != imageData.BoostParam {
				t.Errorf("BoostParam not match: %v", buffEffectRepo.buffEffects.BoostParam)
			}
			t.Logf("BaseParam: %v, BoostParam: %v", buffEffectRepo.buffEffects.BaseParam, buffEffectRepo.buffEffects.BoostParam)
		}
	}
}
