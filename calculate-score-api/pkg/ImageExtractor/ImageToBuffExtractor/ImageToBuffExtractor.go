package Imagetobuffextractor

import (
	"context"
	"os"

	Imagetextextractor "manntera.com/calculate-score-api/pkg/ImageExtractor/ImageTextExtractor"
	"manntera.com/calculate-score-api/pkg/TextReader/BuffSkillReader"
)

func ExtractBuffFromImage(ctx context.Context, file *os.File) ([]BuffSkillReader.BuffSkillParam, error) {
	text, err := Imagetextextractor.ExtractTextFromImage(ctx, file)
	if err != nil {
		return nil, err
	}
	buffSkillParams, err := BuffSkillReader.GetBuffSkillParams(text)
	if err != nil {
		return nil, err
	}
	return buffSkillParams, nil
}
