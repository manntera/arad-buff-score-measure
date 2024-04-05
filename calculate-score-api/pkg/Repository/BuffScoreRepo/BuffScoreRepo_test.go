package BuffScoreRepo

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/Repository/BuffEffectRepo"
	"manntera.com/calculate-score-api/pkg/Repository/BuffSkillRepo"
	"manntera.com/calculate-score-api/pkg/Repository/DetectedTextRepo"
)

func TestBuffScoreRepo(t *testing.T) {
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
		buffEffectRepos := make([]*BuffEffectRepo.BuffEffectRepo, 0)
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
			buffEffectRepo, err := BuffEffectRepo.NewBuffEffectRepoFromDetectedTextRepo(skillRepo, detectedTextRepo)
			if err != nil {
				t.Errorf("Error creating BuffEffectRepo: %v", err)
				continue
			}
			if buffEffectRepo.BuffEffect.BaseParam != imageData.BaseParam {
				t.Errorf("BaseParam not match: %v", buffEffectRepo.BuffEffect.BaseParam)
			}
			if buffEffectRepo.BuffEffect.BoostParam != imageData.BoostParam {
				t.Errorf("BoostParam not match: %v", buffEffectRepo.BuffEffect.BoostParam)
			}
			buffEffectRepos = append(buffEffectRepos, buffEffectRepo)
		}
		buffScoreRepo, err := NewBuffScoreRepo(buffEffectRepos, skillRepo)
		if err != nil {
			t.Errorf("Error creating BuffScoreRepo: %v", err)
			continue
		}
		if buffScoreRepo.BuffScore.Score != testData.Score {
			t.Errorf("Score not match got: %v, want: %v", buffScoreRepo.BuffScore.Score, testData.Score)
		} else {
			log.Printf("Score match got: %v, want: %v", buffScoreRepo.BuffScore.Score, testData.Score)
		}
		if buffScoreRepo.BuffScore.HentaiScore != testData.HentaiScore {
			t.Errorf("HentaiScore not match got: %v, want: %v", buffScoreRepo.BuffScore.HentaiScore, testData.HentaiScore)
		} else {
			log.Printf("HentaiScore match got: %v, want: %v", buffScoreRepo.BuffScore.HentaiScore, testData.HentaiScore)
		}
	}
}
