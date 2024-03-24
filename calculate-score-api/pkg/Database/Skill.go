package Database

type Skill struct {
	ID   int
	Name string
}

var Skills = []Skill{
	{ID: 1, Name: "ラブリーテンポ"},
	{ID: 2, Name: "ステージ"},
	{ID: 3, Name: "有名税"},
	{ID: 4, Name: "アドレナリン"},
	{ID: 5, Name: "特別な物語"},
	{ID: 6, Name: "禁断の呪い"},
	{ID: 7, Name: "マリオネット"},
	{ID: 8, Name: "小悪魔"},
	{ID: 9, Name: "デスティニーパペット"},
	{ID: 10, Name: "終章"},
	{ID: 11, Name: "勇猛の祝福"},
	{ID: 12, Name: "クラクスオブヴィクトリア"},
	{ID: 13, Name: "信実な情熱"},
	{ID: 14, Name: "勇猛のアリア"},
	{ID: 15, Name: "ラウスジアンジェラス"},
	{ID: 16, Name: "グランドクロスクラッシュ"},
	{ID: 17, Name: "栄光の祝福"},
	{ID: 18, Name: "アポカリプス"},
	{ID: 19, Name: "信念のオーラ"},
	{ID: 20, Name: "クロスクラッシュ"},
	{ID: 21, Name: "最後の審判"},
}

func GetSkillFromId(id int) *Skill {
	for _, skill := range Skills {
		if skill.ID == id {
			return &skill
		}
	}
	return nil
}

func GetSkillFromName(name string) *Skill {
	for _, skill := range Skills {
		if skill.Name == name {
			return &skill
		}
	}
	return nil
}
