package ImageProcessing

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

type ImageType int

const (
	PNG ImageType = iota
	JPG ImageType = iota
)

func TrimBuffIconArea(ImageType ImageType, imageFile *os.File) (*os.File, error) {
	sourceImage, err := GetImageFromFile(ImageType, imageFile)
	if err != nil {
		return nil, err
	}
	bounds := sourceImage.Bounds()
	sourceWidth := bounds.Max.X
	sourceHeight := bounds.Max.Y

	halfWidth := sourceWidth / 2
	halfHeight := sourceHeight / 2

	leftX := 0
	rightX := halfWidth
	topY := halfHeight
	bottomY := sourceHeight

	trimmedRect := image.Rect(leftX, topY, rightX, bottomY)
	trimmedImage := image.NewRGBA(trimmedRect)
	draw.Draw(trimmedImage, trimmedImage.Bounds(), sourceImage, image.Pt(leftX, topY), draw.Src)

	tempFile, err := os.CreateTemp("", "trimmed_*.png")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	err = png.Encode(tempFile, trimmedImage)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}

func GetImageFromFile(ImageType ImageType, imageFile *os.File) (image.Image, error) {
	switch ImageType {
	case PNG:
		img, decodeErr := png.Decode(imageFile)
		if decodeErr != nil {
			return nil, decodeErr
		}
		return img, nil
	case JPG:
		img, decodeErr := jpeg.Decode(imageFile)
		if decodeErr != nil {
			return nil, decodeErr
		}
		return img, nil
	default:
		return nil, errors.New("unsupported image type")
	}
}
