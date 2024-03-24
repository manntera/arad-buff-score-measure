package Imagetobuffextractor

import (
	"context"
	"os"

	"manntera.com/calculate-score-api/pkg/Database"
	Imagetextextractor "manntera.com/calculate-score-api/pkg/ImageExtractor/ImageTextExtractor"
	"manntera.com/calculate-score-api/pkg/TextReader/BuffSkillReader"
)

func ExtractBuffFromImage(ctx context.Context, file *os.File) (*Database.BuffSkillParam, error) {
	text, err := Imagetextextractor.ExtractTextFromImage(ctx, file)
	if err != nil {
		return nil, err
	}
	buffSkillParam, err := BuffSkillReader.GetBuffSkillParams(text)
	if err != nil {
		return nil, err
	}
	return &buffSkillParam, nil
}
