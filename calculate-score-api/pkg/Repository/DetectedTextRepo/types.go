package DetectedTextRepo

import (
	"image"

	"manntera.com/calculate-score-api/pkg/NormalizeRect"
)

type DetectedText struct {
	Text          string                      `json:"text"`
	Rect          image.Rectangle             `json:"rect"`
	NormalizeRect NormalizeRect.NormalizeRect `json:"normalizeRect"`
}
