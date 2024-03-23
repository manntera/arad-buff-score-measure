package Database

type BuffSkillParam struct {
	SkillId    int
	BuffParams []BuffParam
}

type BuffParam struct {
	ParamId    int
	ParamValue float64
}
