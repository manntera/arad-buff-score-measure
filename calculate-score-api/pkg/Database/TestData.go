package Database

type ImageData struct {
	ImageFileName string
	SkillId       int
}
type TestData struct {
	JobName          string
	ImageDataList    []ImageData
	SuccessParamList []BuffSkillParam
	Score            int
}

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
		SuccessParamList: []BuffSkillParam{
			{
				SkillId: 1,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 45259,
					},
					{
						ParamId:    2,
						ParamValue: 7871,
					},
				},
			},
			{
				SkillId: 2,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 111539,
					},
				},
			},
			{
				SkillId: 3,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 599,
					},
				},
			},
			{
				SkillId: 4,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 4525,
					},
					{
						ParamId:    2,
						ParamValue: 787,
					},
				},
			},
			{
				SkillId: 5,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 84009,
					},
				},
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
		SuccessParamList: []BuffSkillParam{
			{
				SkillId: 8,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 532,
					},
				},
			},
			{
				SkillId: 6,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 19846,
					},
					{
						ParamId:    2,
						ParamValue: 3823,
					},
				},
			},
			{
				SkillId: 7,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 40710,
					},
				},
			},
			{
				SkillId: 9,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 4961,
					},
					{
						ParamId:    2,
						ParamValue: 955,
					},
				},
			},
			{
				SkillId: 10,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 10991,
					},
				},
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
		SuccessParamList: []BuffSkillParam{
			{
				SkillId: 19,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 397,
					},
				},
			},
			{
				SkillId: 17,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 45930,
					},
					{
						ParamId:    2,
						ParamValue: 8239,
					},
				},
			},
			{
				SkillId: 18,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 84231,
					},
				},
			},
			{
				SkillId: 20,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 53,
					},
				},
			},
			{
				SkillId: 21,
				BuffParams: []BuffParam{
					{
						ParamId:    1,
						ParamValue: 24426,
					},
				},
			},
		},
		Score: 745080,
	},
}

func GetSuccessParamFromSkillId(skillId int) *BuffSkillParam {
	for _, testData := range TestDataList {
		for _, successParam := range testData.SuccessParamList {
			if successParam.SkillId == skillId {
				return &successParam
			}
		}
	}
	return nil
}

func GetSuccessParamFromparamId(BuffSkillParam BuffSkillParam, paramId int) *BuffParam {
	for _, successParam := range BuffSkillParam.BuffParams {
		if successParam.ParamId == paramId {
			return &successParam
		}
	}
	return nil
}
