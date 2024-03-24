package CalculateBuffScore

import (
	"manntera.com/calculate-score-api/pkg/Database"
)

func CalculateBuffScore(buffSkillParams []Database.BuffSkillParam) (int, error) {
	staticParam := 0.0
	ratioParam := 0.0
	for _, buffSkillParam := range buffSkillParams {
		for _, buffParam := range buffSkillParam.BuffParams {
			switch buffParam.ParamId {
			case 1:
				staticParam += buffParam.ParamValue
			case 2:
				ratioParam += buffParam.ParamValue
			}
		}
	}

	staticParam = (staticParam+15000)/250 + 1
	ratioParam = (ratioParam + 2650) / 10
	return int(staticParam * ratioParam), nil
}
