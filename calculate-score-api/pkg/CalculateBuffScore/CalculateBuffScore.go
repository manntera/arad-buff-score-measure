package CalculateBuffScore

import (
	"manntera.com/calculate-score-api/pkg/Database"
	BuffSkillReader "manntera.com/calculate-score-api/pkg/TextReader/BuffSkillReader"
)

func CalculateBuffScore(buffSkillParams []BuffSkillReader.BuffSkillParam) (int, error) {
	staticParam := 0.0
	ratioParam := 0.0
	for _, buffSkillParam := range buffSkillParams {
		for _, buffParam := range buffSkillParam.BuffParams {
			SkillId := buffSkillParam.SkillId

			SkillGenreId := -1
			for _, skill := range Database.Skills {
				if skill.ID == SkillId {
					SkillGenreId = skill.GenreId
					break
				}
			}
			if SkillGenreId == -1 {
				continue
			}

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