package NormalizeRect

type NormalizeRect struct {
	Min NormalizedPoint
	Max NormalizedPoint
}
type NormalizedPoint struct {
	X float64
	Y float64
}

func IsCollisionRect(rect1 NormalizeRect, rect2 NormalizeRect) bool {
	return rect1.Min.X < rect2.Max.X &&
		rect1.Max.X > rect2.Min.X &&
		rect1.Min.Y < rect2.Max.Y &&
		rect1.Max.Y > rect2.Min.Y
}
