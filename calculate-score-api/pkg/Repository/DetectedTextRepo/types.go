package DetectedTextRepo

import (
	"image"

	"manntera.com/calculate-score-api/pkg/NormalizeRect"
)

type DetectedText struct {
	Text          string
	Rect          image.Rectangle
	NormalizeRect NormalizeRect.NormalizeRect
}
