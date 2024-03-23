package BuffParamReader

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BuffParam struct {
	ParamName  string
	ParamValue float64
}

var paramNames = []string{
	"知能",
	"体力",
	"精神力",
	"攻撃速度",
	"移動速度",
	"物理攻撃力",
	"魔法攻撃力",
	"独立攻撃力",
	"的中率",
}

func GetBuffParams(text string, skill string) ([]BuffParam, error) {
	skillPattern := regexp.MustCompile(`(?i)` + skill)

	var skillIndex int = -1
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if skillPattern.MatchString(strings.TrimSpace(line)) {
			skillIndex = i
			break
		}
	}

	if skillIndex == -1 {
		return nil, fmt.Errorf("スキル名 '%s' が見つかりませんでした", skill)
	}

	var buffParams []BuffParam
	for _, line := range lines[skillIndex+1:] {
		line = strings.TrimSpace(line)
		for _, paramName := range paramNames {
			if strings.Contains(line, paramName) && strings.Contains(line, "+") {
				pattern := regexp.MustCompile(`(?i)` + paramName + `.*?([\+\-]?\d+(?:\.\d+)?)`)
				if match := pattern.FindStringSubmatch(line); len(match) > 1 {
					value, _ := strconv.ParseFloat(match[1], 64)
					buffParam := BuffParam{
						ParamName:  paramName,
						ParamValue: value,
					}
					buffParams = append(buffParams, buffParam)
					break
				}
			}
		}
	}

	if len(buffParams) == 0 {
		return nil, fmt.Errorf("スキル '%s' のバフパラメータが見つかりませんでした", skill)
	}

	return buffParams, nil
}
