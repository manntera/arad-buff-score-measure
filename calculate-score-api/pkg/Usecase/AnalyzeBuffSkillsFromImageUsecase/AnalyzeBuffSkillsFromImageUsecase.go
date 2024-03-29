package AnalyzeBuffSkillsFromImageUsecase

import (
	"context"
	"log"
	"os"

	"manntera.com/calculate-score-api/pkg/Database"
	ImageToBuffExtractor "manntera.com/calculate-score-api/pkg/ImageExtractor/ImageToBuffExtractor"
)

func AnalyzeBuffSkillsFromImages(ctx context.Context, files []os.File) ([]Database.BuffSkillParam, error) {
	var buffSkillParams []Database.BuffSkillParam
	for _, file := range files {
		buffSkillParamsFromImage, err := ImageToBuffExtractor.ExtractBuffFromImage(ctx, &file)
		if err != nil {
			log.Printf("画像からスキル情報を抽出できませんでした: %v", err)
			continue
		}
		buffSkillParams = append(buffSkillParams, *buffSkillParamsFromImage)
	}
	return buffSkillParams, nil
}
