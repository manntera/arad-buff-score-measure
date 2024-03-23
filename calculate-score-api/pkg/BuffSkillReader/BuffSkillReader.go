package BuffSkillReader

import "manntera.com/calculate-score-api/pkg/BuffParamReader"

type BuffSkillParam struct {
	SkillName  string
	BuffParams []BuffParamReader.BuffParam
}

var skillNames = []string{
	"ラブリーテンポ",
	"有名税",
	"特別な物語",
	"アドレナリン",
}

func GetBuffSkillParams(text string) ([]BuffSkillParam, error) {
	var buffSkillParams []BuffSkillParam
	for _, skillName := range skillNames {
		buffParams, err := BuffParamReader.GetBuffParams(text, skillName)
		if err != nil {
			continue
		}

		// BuffParamsから重複を排除
		uniqueBuffParams := make([]BuffParamReader.BuffParam, 0)
		paramNames := make(map[string]bool)
		for _, param := range buffParams {
			if !paramNames[param.ParamName] {
				paramNames[param.ParamName] = true
				uniqueBuffParams = append(uniqueBuffParams, param)
			}
		}

		buffSkillParam := BuffSkillParam{
			SkillName:  skillName,
			BuffParams: uniqueBuffParams,
		}
		buffSkillParams = append(buffSkillParams, buffSkillParam)
	}
	return buffSkillParams, nil
}
