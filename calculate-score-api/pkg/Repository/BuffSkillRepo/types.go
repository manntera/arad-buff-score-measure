package BuffSkillRepo

// スキルの情報を保持する構造体
type Skill struct {
	ID           int    // 一意に識別されるID
	Name         string // スキル名
	IsBasePower  bool   // 力/知能のバフを行うか
	IsBoostPower bool   // 魔法/物理/独立攻撃力のバフを行うか
}
