package BuffSkillRepo

import (
	"encoding/json"
	"errors"
	"os"
)

type BuffSkillRepo struct {
	Skills []Skill
}

var _ IBuffSkillRepo = &BuffSkillRepo{}

func NewBuffSkillRepoFromJsonFile(filePath string) (*BuffSkillRepo, error) {

	result := BuffSkillRepo{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var skills []Skill
	err = decoder.Decode(&skills)
	if err != nil {
		return nil, err
	}
	result.Skills = skills

	return &result, nil
}

func (r *BuffSkillRepo) GetSkillList() ([]Skill, error) {
	return r.Skills, nil
}

func (r *BuffSkillRepo) GetSkillFromID(id int) (Skill, error) {
	for _, skill := range r.Skills {
		if skill.ID == id {
			return skill, nil
		}
	}

	return Skill{}, errors.New("Skill not found")
}

func (r *BuffSkillRepo) GetSkillFromName(name string) (Skill, error) {
	for _, skill := range r.Skills {
		if skill.Name == name {
			return skill, nil
		}
	}

	return Skill{}, errors.New("Skill not found")
}
