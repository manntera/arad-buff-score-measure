package BuffParamReader

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"manntera.com/calculate-score-api/pkg/Database"
)

type BuffParam struct {
	ParamId    int
	ParamValue float64
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
		for _, param := range Database.Params {
			paramName := param.Name
			if strings.Contains(line, paramName) && strings.Contains(line, "+") {
				pattern := regexp.MustCompile(`(?i)` + paramName + `.*?([\+\-]?\d+(?:\.\d+)?)`)
				if match := pattern.FindStringSubmatch(line); len(match) > 1 {
					value, _ := strconv.ParseFloat(match[1], 64)
					buffParam := BuffParam{
						ParamId:    param.ID,
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
