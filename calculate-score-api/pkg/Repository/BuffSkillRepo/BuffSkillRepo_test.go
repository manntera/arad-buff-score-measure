package BuffSkillRepo

import (
	"os"
	"testing"
)

func TestBuffSkillRepo(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	dir += "/../../../setting/BuffSkill.json"
	repo, err := NewBuffSkillRepoFromJsonFile(dir)
	if err != nil {
		t.Errorf("Failed to create a new BuffSkillRepo: %v", err)
	}

	skill, err := repo.GetSkillFromID(1)
	if err != nil {
		t.Errorf("Failed to get a skill from ID: %v", err)
	}
	wantSkillName := "ラブリーテンポ"
	if skill.Name != wantSkillName {
		t.Errorf("Expected skill name is %s, but got %s", wantSkillName, skill.Name)
	}
	t.Log(skill)
	wantSlillId := 4
	skill, err = repo.GetSkillFromName("アドレナリン")
	if err != nil {
		t.Errorf("Failed to get a skill from name: %v", err)
	}
	if skill.ID != wantSlillId {
		t.Errorf("Expected skill ID is %d, but got %d", wantSlillId, skill.ID)
	}
	t.Log(skill)
}
