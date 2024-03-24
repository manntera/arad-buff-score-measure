package Imagetobuffextractor

import (
	"context"
	"os"

	"manntera.com/calculate-score-api/pkg/BuffParamTextReader"
	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/ImageExtractor/ImageTextExtractor"
)

// 入力された画像から画像に記載されているスキル情報を抽出し、スキル情報を出力します。
func ExtractBuffFromImage(ctx context.Context, file *os.File) (*Database.BuffSkillParam, error) {
	text, err := ImageTextExtractor.ExtractTextFromImage(ctx, file)
	if err != nil {
		return nil, err
	}
	buffSkillParam, err := BuffParamTextReader.GetBuffParam(text)
	if err != nil {
		return nil, err
	}
	return &buffSkillParam, nil
}
