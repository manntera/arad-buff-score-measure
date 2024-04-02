package DetectedTextRepo

import "manntera.com/calculate-score-api/pkg/NormalizeRect"

type IDetectedTextRepo interface {
	FindLineTextFromKeyword(keyword string, searchRect NormalizeRect.NormalizeRect) ([]*DetectedText, error)
}
