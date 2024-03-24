package BuffSkillReader

import (
	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/TextReader/BuffParamReader"
)

func GetBuffSkillParams(text string) (Database.BuffSkillParam, error) {
	currentSkillParam := Database.BuffSkillParam{}
	for _, skill := range Database.Skills {
		buffParams, err := BuffParamReader.GetBuffParams(text, skill.Name)
		if err != nil {
			continue
		}
		currentSkillParam.SkillId = skill.ID

		uniqueBuffParams := make([]Database.BuffParam, 0)
		paramNames := make(map[int]bool)
		for _, param := range buffParams {
			if !paramNames[param.ParamId] {
				paramNames[param.ParamId] = true
				uniqueBuffParams = append(uniqueBuffParams, param)
			}
		}
		currentSkillParam.BuffParams = uniqueBuffParams
		break
	}
	return currentSkillParam, nil
}
