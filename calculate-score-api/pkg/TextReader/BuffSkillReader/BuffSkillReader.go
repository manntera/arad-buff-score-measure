package BuffSkillReader

import (
	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/TextReader/BuffParamReader"
)

func GetBuffSkillParams(text string) ([]Database.BuffSkillParam, error) {
	var buffSkillParams []Database.BuffSkillParam
	for _, skill := range Database.Skills {
		buffParams, err := BuffParamReader.GetBuffParams(text, skill.Name)
		if err != nil {
			continue
		}

		uniqueBuffParams := make([]Database.BuffParam, 0)
		paramNames := make(map[int]bool)
		for _, param := range buffParams {
			if !paramNames[param.ParamId] {
				paramNames[param.ParamId] = true
				uniqueBuffParams = append(uniqueBuffParams, param)
			}
		}

		buffSkillParam := Database.BuffSkillParam{
			SkillId:    skill.ID,
			BuffParams: uniqueBuffParams,
		}
		buffSkillParams = append(buffSkillParams, buffSkillParam)
	}
	return buffSkillParams, nil
}
