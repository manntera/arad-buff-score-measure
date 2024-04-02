package SamplerImageRepo

import (
	"log"
	"os"
	"testing"

	"manntera.com/calculate-score-api/pkg/Database"
)

func TestSamplerImageRepo(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	dir += "/../../../testdata/"
	for _, testData := range Database.TestDataList {
		log.Printf("【Testing job】 %s", testData.JobName)
		for _, imageData := range testData.ImageDataList {
			log.Printf("【Testing image】 %s", imageData.ImageFileName)
			testImagePath := dir + testData.JobName + "/" + imageData.ImageFileName
			srcFile, err := os.Open(testImagePath)
			if err != nil {
				t.Errorf("Error opening test image: %v", err)
				continue
			}
			repo, err := NewSamplerImageRepoFromFile(srcFile)
			if err != nil {
				t.Errorf("Error creating SamplerImageRepo: %v", err)
				continue
			}
			defer repo.Close()
			repoFile := repo.GetFile()
			repoImage := repo.GetImage()
			repoImageSize := repo.GetImageSize()
			if repoFile == nil {
				t.Errorf("Error getting file from SamplerImageRepo: %v", err)
				continue
			}

			if repoFile.Name() != srcFile.Name() {
				t.Errorf("Error getting file from SamplerImageRepo: %v", err)
				continue
			}

			if repoImage == nil {
				t.Errorf("Error getting image from SamplerImageRepo: %v", err)
				continue
			}
			if repoImageSize.Height != 1080 && repoImageSize.Width != 1920 {
				t.Errorf("Error getting image size from SamplerImageRepo: %v", err)
				continue
			}
		}
	}
}
