package Database

// テスト画像に関する情報を保持する構造体
type ImageData struct {
	ImageFileName string
	SkillId       int
}

// テストデータに関する情報を保持する構造体
type TestData struct {
	JobName              string
	ImageDataList        []ImageData
	ReferenceSkillParams []BuffSkillParam
	Score                int
}

// テストデータ
var TestDataList = []TestData{
	{
		JobName: "muse",
		ImageDataList: []ImageData{
			{
				ImageFileName: "test1.jpg",
				SkillId:       3,
			},
			{
				ImageFileName: "test2.jpg",
				SkillId:       1,
			},
			{
				ImageFileName: "test3.jpg",
				SkillId:       5,
			},
			{
				ImageFileName: "test4.jpg",
				SkillId:       4,
			},
			{
				ImageFileName: "test5.png",
				SkillId:       2,
			},
		},
		ReferenceSkillParams: []BuffSkillParam{
			{
				SkillId:    1,
				BaseParam:  45259,
				BoostParam: 7871,
			},
			{
				SkillId:    2,
				BaseParam:  111539,
				BoostParam: 0,
			},
			{
				SkillId:    3,
				BaseParam:  599,
				BoostParam: 0,
			},
			{
				SkillId:    4,
				BaseParam:  4525,
				BoostParam: 787,
			},
			{
				SkillId:    5,
				BaseParam:  84009,
				BoostParam: 0,
			},
		},
		Score: 1181373,
	},
	{
		JobName: "encha",
		ImageDataList: []ImageData{
			{
				ImageFileName: "test1.png",
				SkillId:       8,
			},
			{
				ImageFileName: "test2.png",
				SkillId:       6,
			},
			{
				ImageFileName: "test3.png",
				SkillId:       7,
			},
			{
				ImageFileName: "test4.png",
				SkillId:       9,
			},
			{
				ImageFileName: "test5.png",
				SkillId:       10,
			},
		},
		ReferenceSkillParams: []BuffSkillParam{
			{
				SkillId:    8,
				BaseParam:  532,
				BoostParam: 0,
			},
			{
				SkillId:    6,
				BaseParam:  19846,
				BoostParam: 3823,
			},
			{
				SkillId:    7,
				BaseParam:  40710,
				BoostParam: 0,
			},
			{
				SkillId:    9,
				BaseParam:  4961,
				BoostParam: 955,
			},
			{
				SkillId:    10,
				BaseParam:  10991,
				BoostParam: 0,
			},
		},
		Score: 274212,
	},
	{
		JobName: "mkuruse",
		ImageDataList: []ImageData{
			{
				ImageFileName: "test1.png",
				SkillId:       19,
			},
			{
				ImageFileName: "test2.png",
				SkillId:       17,
			},
			{
				ImageFileName: "test3.png",
				SkillId:       18,
			},
			{
				ImageFileName: "test4.png",
				SkillId:       20,
			},
			{
				ImageFileName: "test5.png",
				SkillId:       21,
			},
		},
		ReferenceSkillParams: []BuffSkillParam{
			{
				SkillId:   19,
				BaseParam: 397,
			},
			{
				SkillId:    17,
				BaseParam:  45930,
				BoostParam: 8239,
			},
			{
				SkillId:    18,
				BaseParam:  84231,
				BoostParam: 0,
			},
			{
				SkillId:    20,
				BaseParam:  0,
				BoostParam: 53,
			},
			{
				SkillId:    21,
				BaseParam:  24426,
				BoostParam: 0,
			},
		},
		Score: 745080,
	},
	{
		JobName: "wkuruse",
		ImageDataList: []ImageData{
			{
				ImageFileName: "test1.png",
				SkillId:       13,
			},
			{
				ImageFileName: "test2.png",
				SkillId:       11,
			},
			{
				ImageFileName: "test3.png",
				SkillId:       12,
			},
			{
				ImageFileName: "test4.png",
				SkillId:       14,
			},
			{
				ImageFileName: "test5.png",
				SkillId:       15,
			},
			{
				ImageFileName: "test6.png",
				SkillId:       16,
			},
		},
		ReferenceSkillParams: []BuffSkillParam{
			{
				SkillId:    13,
				BaseParam:  509,
				BoostParam: 0,
			},
			{
				SkillId:    11,
				BaseParam:  34904,
				BoostParam: 6384,
			},
			{
				SkillId:    12,
				BaseParam:  52681,
				BoostParam: 0,
			},
			{
				SkillId:    14,
				BaseParam:  5235,
				BoostParam: 957,
			},
			{
				SkillId:    15,
				BaseParam:  14223,
				BoostParam: 0,
			},
			{
				SkillId:    16,
				BaseParam:  288,
				BoostParam: 0,
			},
		},
		Score: 491916,
	},
}

// スキルIDからテストデータを取得する
func GetSuccessParamFromSkillId(skillId int) *BuffSkillParam {
	for _, testData := range TestDataList {
		for _, successParam := range testData.ReferenceSkillParams {
			if successParam.SkillId == skillId {
				return &successParam
			}
		}
	}
	return nil
}
