package BuffScoreRepo

import (
	"manntera.com/calculate-score-api/pkg/Repository/BuffEffectRepo"
	"manntera.com/calculate-score-api/pkg/Repository/BuffSkillRepo"
)

type BuffScoreRepo struct {
	BuffScore *BuffScore
}

var _ IBuffScoreRepo = &BuffScoreRepo{}

func NewBuffScoreRepo(buffEffectRepos []*BuffEffectRepo.BuffEffectRepo, buffSkillRepo *BuffSkillRepo.BuffSkillRepo) (*BuffScoreRepo, error) {
	result := &BuffScoreRepo{}
	BuffScore := &BuffScore{}

	henaiSkill, err := buffSkillRepo.GetSkillFromName("禁断の呪い")
	if err != nil {
		return nil, err
	}

	var baseParam float32 = 0
	var boostParam float32 = 0
	var henaiBaseScore float32 = 0
	var henaiBoostScore float32 = 0

	for _, buffEffectRepo := range buffEffectRepos {
		baseParam += float32(buffEffectRepo.BuffEffect.BaseParam)
		boostParam += float32(buffEffectRepo.BuffEffect.BoostParam)

		// 偏愛スキルのスコア計算
		if buffEffectRepo.BuffEffect.SkillId == henaiSkill.ID {
			henaiBaseScore += (float32(buffEffectRepo.BuffEffect.BaseParam) * 1.15)
			henaiBoostScore += (float32(buffEffectRepo.BuffEffect.BoostParam) * 1.15)
		} else {
			henaiBaseScore += float32(buffEffectRepo.BuffEffect.BaseParam)
			henaiBoostScore += float32(buffEffectRepo.BuffEffect.BoostParam)
		}
	}

	baseParam = (baseParam+15000.0)/250.0 + 1.0
	boostParam = (boostParam + 2650.0) / 10.0
	BuffScore.Score = int(baseParam * boostParam)

	henaiBaseScore = (henaiBaseScore+15000.0)/250.0 + 1.0
	henaiBoostScore = (henaiBoostScore + 2650.0) / 10.0
	BuffScore.HentaiScore = int(henaiBaseScore * henaiBoostScore)

	result.BuffScore = BuffScore

	return result, nil
}
