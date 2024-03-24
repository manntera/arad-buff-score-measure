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
