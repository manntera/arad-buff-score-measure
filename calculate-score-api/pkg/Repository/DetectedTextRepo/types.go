package DetectedTextRepo

import (
	"image"

	"manntera.com/calculate-score-api/pkg/NormalizeRect"
)

type DetectedText struct {
	text          string
	rect          image.Rectangle
	normalizeRect NormalizeRect.NormalizeRect
}
