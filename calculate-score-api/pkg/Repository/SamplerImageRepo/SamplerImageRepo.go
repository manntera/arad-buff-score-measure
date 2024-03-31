package SamplerImageRepo

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
)

type SamplerImageRepo struct {
	samplerImage *SamplerImage
}

var _ IImageSamplerRepo = &SamplerImageRepo{}

func NewSamplerImageRepoFromFileHeader(fileHeader *multipart.FileHeader) (*SamplerImageRepo, error) {
	result := &SamplerImageRepo{}

	SamplerImage := SamplerImage{
		srcFile:  nil,
		srcImage: nil,
	}

	file, err := createFileFromFileHeader(fileHeader)
	if err != nil {
		return nil, err
	}

	image, err := getImageFromFile(PNG, file)
	if err != nil {
		return nil, err
	}

	SamplerImage.srcImage = &image
	SamplerImage.srcFile = file
	result.samplerImage = &SamplerImage

	return result, nil
}

func (s *SamplerImageRepo) Close() error {
	sampler := s.samplerImage

	if sampler.srcFile != nil {
		return sampler.srcFile.Close()
	}
	sampler.srcImage = nil

	return nil
}

func (s *SamplerImageRepo) GetFile() *os.File {
	return s.samplerImage.srcFile
}

func (s *SamplerImageRepo) GetImage() *image.Image {
	return s.samplerImage.srcImage
}

func (s *SamplerImageRepo) GetImageSize() *Size {
	img := *s.GetImage()
	bounds := img.Bounds()
	return &Size{
		Width:  bounds.Dx(),
		Height: bounds.Dy(),
	}
}

func createFileFromFileHeader(fileHeader *multipart.FileHeader) (*os.File, error) {
	src, _ := fileHeader.Open()
	defer src.Close()

	tempFile, err := os.CreateTemp("", "image")
	if err != nil {
		return nil, err
	}

	_, err = tempFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}

func getImageFromFile(ImageType ImageType, imageFile *os.File) (image.Image, error) {
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
