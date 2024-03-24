package CalculateBuffScore

import (
	"manntera.com/calculate-score-api/pkg/Database"
)

// バフスキルのパラメーターを総合して、バフスキルのスコアを計算します。
func CalculateBuffScore(buffSkillParams []Database.BuffSkillParam) (int, error) {
	baseParam := 0.0
	boostParam := 0.0
	for _, buffSkillParam := range buffSkillParams {
		baseParam += float64(buffSkillParam.BaseParam)
		boostParam += float64(buffSkillParam.BoostParam)
	}

	baseParam = (baseParam+15000.0)/250.0 + 1.0
	boostParam = (boostParam + 2650.0) / 10.0
	return int(baseParam * boostParam), nil
}
