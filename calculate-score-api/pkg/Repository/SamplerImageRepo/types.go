package SamplerImageRepo

import (
	"image"
	"os"
)

type SamplerImage struct {
	Width    int
	Height   int
	SrcFile  *os.File
	SrcImage *image.Image
}

type ImageType int

const (
	PNG ImageType = iota
	JPG ImageType = iota
)
