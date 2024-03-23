package AnalyzeBuffSkillsFromImageUsecase

import (
	"context"
	"os"

	ImageToBuffExtractor "manntera.com/calculate-score-api/pkg/ImageExtractor/ImageToBuffExtractor"
	BuffSkillReader "manntera.com/calculate-score-api/pkg/TextReader/BuffSkillReader"
)

func AnalyzeBuffSkillsFromImages(ctx context.Context, files []os.File) ([]BuffSkillReader.BuffSkillParam, error) {
	var buffSkillParams []BuffSkillReader.BuffSkillParam
	for _, file := range files {
		buffSkillParamsFromImage, err := ImageToBuffExtractor.ExtractBuffFromImage(ctx, &file)
		if err != nil {
			return nil, err
		}
		buffSkillParams = append(buffSkillParams, buffSkillParamsFromImage...)
	}
	return buffSkillParams, nil
}
