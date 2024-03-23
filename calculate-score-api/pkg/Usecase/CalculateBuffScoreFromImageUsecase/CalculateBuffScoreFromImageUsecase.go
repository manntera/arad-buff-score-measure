package CalculateBuffScoreFromImageUsecase

import (
	"context"
	"os"

	"manntera.com/calculate-score-api/pkg/CalculateBuffScore"
	"manntera.com/calculate-score-api/pkg/Usecase/AnalyzeBuffSkillsFromImageUsecase"
)

func CalculateBuffScoreFromImage(ctx context.Context, images []os.File) (int, error) {
	buffSkillParams, err := AnalyzeBuffSkillsFromImageUsecase.AnalyzeBuffSkillsFromImages(ctx, images)
	if err != nil {
		return -1, err
	}

	score, err := CalculateBuffScore.CalculateBuffScore(buffSkillParams)
	if err != nil {
		return -1, err
	}

	return score, nil
}
