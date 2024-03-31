package Database

import "fmt"

// テスト画像に関する情報を保持する構造体
type ImageData struct {
	ImageFileName string // テスト画像ファイル名
	SkillId       int    // 画像で表示されているスキルのID(SkillsのIdに対応する)
	BaseParam     int    // 力/知能の上昇値
	BoostParam    int    // 物理/魔法/独立攻撃力の上昇値
}

// テストデータに関する情報を保持する構造体
type TestData struct {
	JobName       string      // 職業名
	ImageDataList []ImageData // テスト使用する画像リスト
	Score         int         // 期待されるスコア
}

// テストデータ
var TestDataList = []TestData{
	{
		JobName: "muse",
		ImageDataList: []ImageData{
			{
				ImageFileName: "test1.png",
				SkillId:       3, // 有名税
				BaseParam:     599,
				BoostParam:    0,
			},
			{
				ImageFileName: "test2.png",
				SkillId:       1, // ラブリーテンポ
				BaseParam:     45259,
				BoostParam:    7871,
			},
			{
				ImageFileName: "test3.png",
				SkillId:       5, // 特別な物語
				BaseParam:     84009,
				BoostParam:    0,
			},
			{
				ImageFileName: "test4.png",
				SkillId:       4, // アドレナリン
				BaseParam:     4525,
				BoostParam:    787,
			},
			{
				ImageFileName: "test5.png",
				SkillId:       2, // ステージ
				BaseParam:     111539,
				BoostParam:    0,
			},
		},
		Score: 1181373,
	},
	{
		JobName: "encha",
		ImageDataList: []ImageData{
			{
				ImageFileName: "test1.png",
				SkillId:       8, // 小悪魔
				BaseParam:     532,
				BoostParam:    0,
			},
			{
				ImageFileName: "test2.png",
				SkillId:       6, // 禁断の呪い
				BaseParam:     19846,
				BoostParam:    3823,
			},
			{
				ImageFileName: "test3.png",
				SkillId:       7, // マリオネット
				BaseParam:     40710,
				BoostParam:    0,
			},
			{
				ImageFileName: "test4.png",
				SkillId:       9, // デスティニーパペット
				BaseParam:     4961,
				BoostParam:    955,
			},
			{
				ImageFileName: "test5.png",
				SkillId:       10, // 終章
				BaseParam:     10991,
				BoostParam:    0,
			},
		},
		Score: 274212,
	},
	{
		JobName: "mkuruse",
		ImageDataList: []ImageData{
			{
				ImageFileName: "test1.png",
				SkillId:       19, // 信念のオーラ
				BaseParam:     397,
				BoostParam:    0,
			},
			{
				ImageFileName: "test2.png",
				SkillId:       17, // 栄光の祝福
				BaseParam:     45930,
				BoostParam:    8239,
			},
			{
				ImageFileName: "test3.png",
				SkillId:       18, // アポカリプス
				BaseParam:     84231,
				BoostParam:    0,
			},
			{
				ImageFileName: "test4.png",
				SkillId:       20, // クロスクラッシュ
				BaseParam:     0,
				BoostParam:    53,
			},
			{
				ImageFileName: "test5.png",
				SkillId:       21, // 最後の審判
				BaseParam:     24426,
				BoostParam:    0,
			},
		},
		Score: 745080,
	},
	{
		JobName: "wkuruse",
		ImageDataList: []ImageData{
			{
				ImageFileName: "test1.png",
				SkillId:       13, // 信実な情熱
				BaseParam:     509,
				BoostParam:    0,
			},
			{
				ImageFileName: "test2.png",
				SkillId:       11, // 勇猛の祝福
				BaseParam:     34904,
				BoostParam:    6384,
			},
			{
				ImageFileName: "test3.png",
				SkillId:       12, // クラクス
				BaseParam:     52681,
				BoostParam:    0,
			},
			{
				ImageFileName: "test4.png",
				SkillId:       14, // 勇猛のアリア
				BaseParam:     5235,
				BoostParam:    957,
			},
			{
				ImageFileName: "test5.png",
				SkillId:       15, // ラウス
				BaseParam:     14223,
				BoostParam:    0,
			},
			{
				ImageFileName: "test6.png",
				SkillId:       16, // グランドクロスクラッシュ
				BaseParam:     288,
				BoostParam:    0,
			},
		},
		Score: 491916,
	},
}

// スキルIDからテストデータを取得する
func GetImageDataFromSkillId(skillId int) (ImageData, error) {
	for _, testData := range TestDataList {
		for _, imageData := range testData.ImageDataList {
			if imageData.SkillId == skillId {
				return imageData, nil
			}
		}
	}
	return ImageData{}, fmt.Errorf("SkillId %d is not found", skillId)
}
