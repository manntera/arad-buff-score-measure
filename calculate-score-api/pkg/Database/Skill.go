package Database

type Skill struct {
	ID      int
	Name    string
	GenreId int
}

var Skills = []Skill{
	{ID: 1, Name: "ラブリーテンポ", GenreId: 1},
	{ID: 2, Name: "ステージ", GenreId: 2},
	{ID: 3, Name: "有名税", GenreId: 3},
	{ID: 4, Name: "アドレナリン", GenreId: 4},
	{ID: 5, Name: "特別な物語", GenreId: 5},
	{ID: 6, Name: "禁断の呪い", GenreId: 1},
	{ID: 7, Name: "マリオネット", GenreId: 2},
	{ID: 8, Name: "小悪魔", GenreId: 3},
	{ID: 9, Name: "デスティニーパペット", GenreId: 4},
	{ID: 10, Name: "終章", GenreId: 5},
	{ID: 11, Name: "勇猛の祝福", GenreId: 1},
	{ID: 12, Name: "クラクスオブヴィクトリア", GenreId: 2},
	{ID: 13, Name: "信実な情熱", GenreId: 3},
	{ID: 14, Name: "勇猛のアリア", GenreId: 4},
	{ID: 15, Name: "ラウスジアンジェラス", GenreId: 5},
	{ID: 16, Name: "グランドクロスクラッシュ", GenreId: 6},
	{ID: 17, Name: "栄光の祝福", GenreId: 1},
	{ID: 18, Name: "アポカリプス", GenreId: 2},
	{ID: 19, Name: "信念のオーラ", GenreId: 3},
	{ID: 20, Name: "クロスクラッシュ", GenreId: 4},
	{ID: 21, Name: "最後の審判", GenreId: 5},
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
