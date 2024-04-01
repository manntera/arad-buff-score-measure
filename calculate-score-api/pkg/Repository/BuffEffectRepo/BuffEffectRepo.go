package BuffEffectRepo

import (
	"fmt"
	"strings"

	"manntera.com/calculate-score-api/pkg/NormalizeRect"
	"manntera.com/calculate-score-api/pkg/Repository/BuffSkillRepo"
	"manntera.com/calculate-score-api/pkg/Repository/DetectedTextRepo"
)

type BuffEffectRepo struct {
	buffEffects []*BuffEffect
}

var _ IBuffEffectRepo = &BuffEffectRepo{}

func NewBuffEffectRepoFromDetectedTextRepo(buffSkillRepo BuffSkillRepo.IBuffSkillRepo, detectedTextRepo DetectedTextRepo.IDetectedTextRepo) (*BuffEffectRepo, error) {
	result := BuffEffectRepo{}
	skills, err := buffSkillRepo.GetSkillList()
	if err != nil {
		return nil, err
	}
	searchTargetRects := NormalizeRect.NormalizeRect{
		Min: NormalizeRect.NormalizedPoint{X: 0.0, Y: 0.5},
		Max: NormalizeRect.NormalizedPoint{X: 0.5, Y: 1.0},
	}
	var findSkill *DetectedTextRepo.DetectedText = nil
	for _, skill := range skills {
		detectedSkillTexts, err := detectedTextRepo.FindLineTextFromKeyword(skill.Name, searchTargetRects)
		if err != nil {
			continue
		}
		if len(detectedSkillTexts) == 0 {
			continue
		}
		for _, detectedSkillText := range detectedSkillTexts {
			if detectedSkillText.Rect.Min.Y > findSkill.Rect.Min.Y {
				findSkill = detectedSkillText
			}
		}
	}
	if findSkill == nil {
		return nil, fmt.Errorf("skill not found")
	}
	tempintBuffs, err := detectedTextRepo.FindLineTextFromKeyword("知能", searchTargetRects)
	if err != nil {
		return nil, err
	}
	IntBuffs := []*DetectedTextRepo.DetectedText{}
	for _, intBuff := range tempintBuffs {
		if !strings.Contains(intBuff.Text, "適用") {
			IntBuffs = append(IntBuffs, intBuff)
		}
	}

	tempPowBuffs, err := detectedTextRepo.FindLineTextFromKeyword("力", searchTargetRects)
	if err != nil {
		return nil, err
	}
	PowBuffs := []*DetectedTextRepo.DetectedText{}
	for _, powBuff := range tempPowBuffs {
		if !strings.Contains(powBuff.Text, "物理") &&
			!strings.Contains(powBuff.Text, "魔法") &&
			!strings.Contains(powBuff.Text, "攻撃") &&
			!strings.Contains(powBuff.Text, "適用") &&
			!strings.Contains(powBuff.Text, "独立") {
			PowBuffs = append(PowBuffs, powBuff)
		}
	}

	// IntBuffとPowBuffの中で一番信憑性が高いやつを採用する
	// 次にBurstBuffを検出する

	return &result, nil
}

func (repo *BuffEffectRepo) CalculateBuffScore() (int, error) {
	baseParam := 0.0
	boostParam := 0.0
	for _, buffEffect := range repo.buffEffects {
		baseParam += float64(buffEffect.BaseParam)
		boostParam += float64(buffEffect.BoostParam)
	}

	baseParam = (baseParam+15000.0)/250.0 + 1.0
	boostParam = (boostParam + 2650.0) / 10.0
	return int(baseParam * boostParam), nil
}
