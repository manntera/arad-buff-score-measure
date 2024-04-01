package DetectedTextRepo

import "image"

type DetectedText struct {
	text          string
	rect          image.Rectangle
	normalizeRect NormalizeRect
}
type NormalizeRect struct {
	Min NormalizedPoint
	Max NormalizedPoint
}
type NormalizedPoint struct {
	X float64
	Y float64
}
