package BuffEffectRepo

type BuffEffect struct {
	SkillId    int // スキルを一意に識別するID
	BaseParam  int // 力/知能の上昇値
	BoostParam int // 物理/魔法/独立攻撃力の上昇値
}
