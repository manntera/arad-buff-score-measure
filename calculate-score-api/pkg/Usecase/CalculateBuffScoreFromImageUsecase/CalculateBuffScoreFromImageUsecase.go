package CalculateBuffScoreFromImageUsecase

import (
	"context"
	"os"

	"manntera.com/calculate-score-api/pkg/CalculateBuffScore"
	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/Usecase/AnalyzeBuffSkillsFromImageUsecase"
)

func CalculateBuffScoreFromImage(ctx context.Context, images []os.File) (int, []Database.BuffSkillParam, error) {
	buffSkillParams, err := AnalyzeBuffSkillsFromImageUsecase.AnalyzeBuffSkillsFromImages(ctx, images)
	if err != nil {
		return -1, nil, err
	}

	score, err := CalculateBuffScore.CalculateBuffScore(buffSkillParams)
	if err != nil {
		return -1, nil, err
	}

	return score, buffSkillParams, nil
}
