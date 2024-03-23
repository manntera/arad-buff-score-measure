package Database

type Skill struct {
	ID      int
	Name    string
	GenreId int
}

var Skills = []Skill{
	{ID: 1, Name: "ラブリーテンポ", GenreId: 1},
	{ID: 2, Name: "ステージ", GenreId: 2},
	{ID: 3, Name: "有名税", GenreId: 1},
	{ID: 4, Name: "アドレナリン", GenreId: 1},
	{ID: 5, Name: "特別な物語", GenreId: 1},
}
