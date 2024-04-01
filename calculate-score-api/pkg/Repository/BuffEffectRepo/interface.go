package BuffEffectRepo

type IBuffEffectRepo interface {
	CalculateBuffScore() (int, error)
}
