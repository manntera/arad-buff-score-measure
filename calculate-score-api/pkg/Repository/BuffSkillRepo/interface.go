package BuffSkillRepo

type IBuffSkillRepo interface {
	GetSkillFromID(id int) (*Skill, error)
	GetSkillFromName(name string) (*Skill, error)
}
