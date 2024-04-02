package SamplerImageRepo

import (
	"image"
	"os"
)

type SamplerImage struct {
	srcFile  *os.File
	srcImage *image.Image
}

type ImageType int

type Size struct {
	Width  int
	Height int
}

const (
	PNG ImageType = iota
	JPG ImageType = iota
)
