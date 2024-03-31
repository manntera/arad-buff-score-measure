package SamplerImageRepo

import (
	"image"
	"os"
)

type IImageSamplerRepo interface {
	Close() error
	GetFile() *os.File
	GetImage() *image.Image
	GetImageSize() *Size
}
