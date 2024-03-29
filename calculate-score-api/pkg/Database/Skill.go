package Database

import "fmt"

// スキルの情報を保持する構造体
type Skill struct {
	ID           int    // 一意に識別されるID
	Name         string // スキル名
	IsBasePower  bool   // 力/知能のバフを行うか
	IsBoostPower bool   // 魔法/物理/独立攻撃力のバフを行うか
}

var Skills = []Skill{
	{ID: 1, Name: "ラブリーテンポ", IsBasePower: true, IsBoostPower: true},
	{ID: 2, Name: "ステージ", IsBasePower: true, IsBoostPower: false},
	{ID: 3, Name: "有名税", IsBasePower: true, IsBoostPower: false},
	{ID: 4, Name: "アドレナリン", IsBasePower: true, IsBoostPower: true},
	{ID: 5, Name: "FINALE", IsBasePower: true, IsBoostPower: false},
	{ID: 6, Name: "禁断の呪い", IsBasePower: true, IsBoostPower: true},
	{ID: 7, Name: "マリオネット", IsBasePower: true, IsBoostPower: false},
	{ID: 8, Name: "小悪魔", IsBasePower: true, IsBoostPower: false},
	{ID: 9, Name: "デスティニーパペット", IsBasePower: true, IsBoostPower: true},
	{ID: 10, Name: "終章", IsBasePower: true, IsBoostPower: false},
	{ID: 11, Name: "勇猛の祝福", IsBasePower: true, IsBoostPower: true},
	{ID: 12, Name: "クラクス", IsBasePower: true, IsBoostPower: false},
	{ID: 13, Name: "情熱", IsBasePower: true, IsBoostPower: false},
	{ID: 14, Name: "勇猛のアリア", IsBasePower: true, IsBoostPower: true},
	{ID: 15, Name: "ラウス", IsBasePower: true, IsBoostPower: false},
	{ID: 16, Name: "グランドクロスクラッシュ", IsBasePower: true, IsBoostPower: false},
	{ID: 17, Name: "栄光の祝福", IsBasePower: true, IsBoostPower: true},
	{ID: 18, Name: "アポカリプス", IsBasePower: true, IsBoostPower: false},
	{ID: 19, Name: "信念のオーラ", IsBasePower: true, IsBoostPower: false},
	{ID: 20, Name: "クロスクラッシュ", IsBasePower: false, IsBoostPower: true},
	{ID: 21, Name: "最後の審判", IsBasePower: true, IsBoostPower: false},
}

// IDからスキル情報を取得する
func GetSkillFromId(id int) (Skill, error) {
	for _, skill := range Skills {
		if skill.ID == id {
			return skill, nil
		}
	}
	return Skill{}, fmt.Errorf("スキルが見つかりませんでした")
}

// スキル名からスキル情報を取得する
func GetSkillFromName(name string) (Skill, error) {
	for _, skill := range Skills {
		if skill.Name == name {
			return skill, nil
		}
	}
	return Skill{}, fmt.Errorf("スキルが見つかりませんでした")
}
