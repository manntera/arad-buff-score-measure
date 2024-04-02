package BuffEffectRepo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"manntera.com/calculate-score-api/pkg/NormalizeRect"
	"manntera.com/calculate-score-api/pkg/Repository/BuffSkillRepo"
	"manntera.com/calculate-score-api/pkg/Repository/DetectedTextRepo"
)

type BuffEffectRepo struct {
	buffEffects *BuffEffect
}

var _ IBuffEffectRepo = &BuffEffectRepo{}

const (
	searchMinX = 0.0
	searchMinY = 0.5
	searchMaxX = 0.5
	searchMaxY = 1.0
)

func NewBuffEffectRepoFromDetectedTextRepo(buffSkillRepo BuffSkillRepo.IBuffSkillRepo, detectedTextRepo DetectedTextRepo.IDetectedTextRepo) (*BuffEffectRepo, error) {
	result := BuffEffectRepo{
		buffEffects: &BuffEffect{},
	}

	skillName, err := findSkillName(buffSkillRepo, detectedTextRepo)
	if err != nil {
		return nil, err
	}

	skill, err := buffSkillRepo.GetSkillFromName(skillName)
	if err != nil {
		return nil, err
	}
	result.buffEffects.SkillId = skill.ID

	if skill.IsBasePower {
		result.buffEffects.BaseParam, err = findParam(detectedTextRepo, []string{"知能", "力"}, []string{"物理", "魔法", "攻撃", "適用", "独立"})
		if err != nil {
			return nil, err
		}
	}

	if skill.IsBoostPower {
		result.buffEffects.BoostParam, err = findParam(detectedTextRepo, []string{"魔法", "物理", "独立"}, []string{"適用"})
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func findSkillName(buffSkillRepo BuffSkillRepo.IBuffSkillRepo, detectedTextRepo DetectedTextRepo.IDetectedTextRepo) (string, error) {
	skills, err := buffSkillRepo.GetSkillList()
	if err != nil {
		return "", err
	}

	searchTargetRects := NormalizeRect.NormalizeRect{
		Min: NormalizeRect.NormalizedPoint{X: searchMinX, Y: searchMinY},
		Max: NormalizeRect.NormalizedPoint{X: searchMaxX, Y: searchMaxY},
	}

	var findSkill DetectedTextRepo.DetectedText
	var skillName string
	for _, skill := range skills {
		detectedSkillTexts, err := detectedTextRepo.FindLineTextFromKeyword(skill.Name, searchTargetRects)
		if err != nil || len(detectedSkillTexts) == 0 {
			continue
		}

		for _, detectedSkillText := range detectedSkillTexts {
			if detectedSkillText.Rect.Min.Y > findSkill.Rect.Min.Y {
				findSkill = *detectedSkillText
				skillName = skill.Name
			}
		}
	}

	if findSkill.Text == "" {
		return "", fmt.Errorf("skill not found")
	}
	return skillName, nil
}

func findParam(detectedTextRepo DetectedTextRepo.IDetectedTextRepo, includeKeywords, excludeKeywords []string) (int, error) {
	searchTargetRects := NormalizeRect.NormalizeRect{
		Min: NormalizeRect.NormalizedPoint{X: searchMinX, Y: searchMinY},
		Max: NormalizeRect.NormalizedPoint{X: searchMaxX, Y: searchMaxY},
	}

	var tempBuffs []*DetectedTextRepo.DetectedText
	for _, keyword := range includeKeywords {
		tempBuff, _ := detectedTextRepo.FindLineTextFromKeyword(keyword, searchTargetRects)
		tempBuffs = append(tempBuffs, tempBuff...)
	}

	var maxParam int
	for _, tempBuff := range tempBuffs {
		text := tempBuff.Text
		if !containsAny(text, excludeKeywords) {
			for _, keyword := range includeKeywords {
				tempParam, err := getNumberFromLine(text, keyword)
				if err == nil && tempParam > maxParam {
					maxParam = tempParam
					break
				}
			}
		}
	}

	return maxParam, nil
}

func containsAny(text string, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	return false
}

func getNumberFromLine(line string, paramName string) (int, error) {
	pattern := regexp.MustCompile(`(?i)` + paramName + `.*?([\+\-]?\d+(?:\.\d+)?)`)
	if match := pattern.FindStringSubmatch(line); len(match) > 1 {
		number, err := strconv.Atoi(match[1])
		if err == nil {
			return number, nil
		}
		return 0, fmt.Errorf("failed to convert number: %v", err)
	}
	return 0, fmt.Errorf("parameter not found: %s", paramName)
}
