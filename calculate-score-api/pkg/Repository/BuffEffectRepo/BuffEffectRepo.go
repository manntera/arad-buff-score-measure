package BuffEffectRepo

type BuffEffectRepo struct {
	buffEffects []*BuffEffect
}

var _ IBuffEffectRepo = &BuffEffectRepo{}

func NewBuffEffectRepo(buffEffects []*BuffEffect) (*BuffEffectRepo, error) {
	result := BuffEffectRepo{}
	result.buffEffects = buffEffects
	return &result, nil
}

func (repo *BuffEffectRepo) CalculateBuffScore() (int, error) {
	baseParam := 0.0
	boostParam := 0.0
	for _, buffEffect := range repo.buffEffects {
		baseParam += float64(buffEffect.BaseParam)
		boostParam += float64(buffEffect.BoostParam)
	}

	baseParam = (baseParam+15000.0)/250.0 + 1.0
	boostParam = (boostParam + 2650.0) / 10.0
	return int(baseParam * boostParam), nil
}
