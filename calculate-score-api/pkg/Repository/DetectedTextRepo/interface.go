package DetectedTextRepo

import "image"

type IDetectedTextRepo interface {
	FindLineTextFromKeyword(keyword string, searchRect image.Rectangle) ([]*DetectedText, error)
}
