package BuffParamTextReader

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"manntera.com/calculate-score-api/pkg/Database"
)

// テキストからスキルを検出し、そのスキルのバフパラメータを取得します。
func GetBuffParam(text string) (Database.BuffSkillParam, error) {
	detectedSkillName := ""
	detectedSkillNameLineIndex := -1
	lines := strings.Split(text, "\n")
	result := Database.BuffSkillParam{}
	for _, skill := range Database.Skills {
		for index, line := range lines {
			if strings.Contains(line, skill.Name) {
				log.Default().Println("【Found skill name】 " + line)
				detectedSkillName = skill.Name
				detectedSkillNameLineIndex = index
				break
			}
		}
	}
	if detectedSkillName == "" {
		return Database.BuffSkillParam{}, fmt.Errorf("スキルが見つかりませんでした")
	}
	log.Printf("検出されたスキル名: %s", detectedSkillName)

	currentSkill, getSkillErr := Database.GetSkillFromName(detectedSkillName)
	if getSkillErr != nil {
		return Database.BuffSkillParam{}, getSkillErr
	}
	result.SkillId = currentSkill.ID

	for _, line := range lines[detectedSkillNameLineIndex+1:] {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "適用") {
			continue
		}
		if strings.Contains(line, "攻撃力") {
			param, err := getNumberFromLine(line, "攻撃力")
			if err != nil {
				continue
			}
			if result.BoostParam == 0 {
				result.BoostParam = param
			}

			continue
		}
		if strings.Contains(line, "物理") {
			param, err := getNumberFromLine(line, "物理")
			if err != nil {
				continue
			}
			if result.BoostParam == 0 {
				result.BoostParam = param
			}
			continue
		}
		if strings.Contains(line, "魔法") {
			param, err := getNumberFromLine(line, "魔法")
			if err != nil {
				continue
			}
			if result.BoostParam == 0 {
				result.BoostParam = param
			}
			continue
		}
		if strings.Contains(line, "独立") {
			param, err := getNumberFromLine(line, "独立")
			if err != nil {
				continue
			}
			if result.BoostParam == 0 {
				result.BoostParam = param
			}
			continue
		}
		if strings.Contains(line, "知能") {
			param, err := getNumberFromLine(line, "知能")
			if err != nil {
				continue
			}
			if result.BaseParam == 0 {
				result.BaseParam = param
			}
			continue
		}
		if strings.Contains(line, "力") {
			if strings.Contains(line, "体力") {
				continue
			}
			if strings.Contains(line, "精神力") {
				continue
			}
			if strings.Contains(line, "能力") {
				continue
			}
			if strings.Contains(line, "物理") {
				continue
			}
			if strings.Contains(line, "魔法") {
				continue
			}
			if strings.Contains(line, "独立") {
				continue
			}
			param, err := getNumberFromLine(line, "力")
			if err != nil {
				continue
			}
			if result.BaseParam == 0 {
				result.BaseParam = param
			}
			continue
		}
	}
	return result, nil
}

func getNumberFromLine(line string, paramName string) (int, error) {
	pattern := regexp.MustCompile(`(?i)` + paramName + `.*?([\+\-]?\d+(?:\.\d+)?)`)
	if match := pattern.FindStringSubmatch(line); len(match) > 1 {
		if number, err := strconv.Atoi(match[1]); err == nil {
			return number, nil
		}
		return 0, fmt.Errorf("数値の変換に失敗しました")
	}
	return 0, fmt.Errorf("パラメータが見つかりませんでした")
}
