package BuffSkillReader

import (
	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/TextReader/BuffParamReader"
)

type BuffSkillParam struct {
	SkillId    int
	BuffParams []BuffParamReader.BuffParam
}

func GetBuffSkillParams(text string) ([]BuffSkillParam, error) {
	var buffSkillParams []BuffSkillParam
	for _, skill := range Database.Skills {
		buffParams, err := BuffParamReader.GetBuffParams(text, skill.Name)
		if err != nil {
			continue
		}

		uniqueBuffParams := make([]BuffParamReader.BuffParam, 0)
		paramNames := make(map[int]bool)
		for _, param := range buffParams {
			if !paramNames[param.ParamId] {
				paramNames[param.ParamId] = true
				uniqueBuffParams = append(uniqueBuffParams, param)
			}
		}

		buffSkillParam := BuffSkillParam{
			SkillId:    skill.ID,
			BuffParams: uniqueBuffParams,
		}
		buffSkillParams = append(buffSkillParams, buffSkillParam)
	}
	return buffSkillParams, nil
}
