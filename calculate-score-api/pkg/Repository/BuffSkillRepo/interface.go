package BuffSkillRepo

type IBuffSkillRepo interface {
	GetSkillList() ([]Skill, error)
	GetSkillFromID(id int) (Skill, error)
	GetSkillFromName(name string) (Skill, error)
}
