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
	SamplerImage *SamplerImage
}

var _ IImageSamplerRepo = &SamplerImageRepo{}

func NewSamplerImageRepoFromFileHeader(fileHeader *multipart.FileHeader) (*SamplerImageRepo, error) {
	result := &SamplerImageRepo{}

	SamplerImage := SamplerImage{
		SrcFile:  nil,
		SrcImage: nil,
	}

	file, err := CreateFileFromFileHeader(fileHeader)
	if err != nil {
		return nil, err
	}

	image, err := GetImageFromFile(PNG, file)
	if err != nil {
		return nil, err
	}

	SamplerImage.SrcImage = &image
	SamplerImage.SrcFile = file
	result.SamplerImage = &SamplerImage

	return result, nil
}

func (s *SamplerImageRepo) Close() error {
	sampler := s.SamplerImage

	if sampler.SrcFile != nil {
		return sampler.SrcFile.Close()
	}
	sampler.SrcImage = nil

	return nil
}

func CreateFileFromFileHeader(fileHeader *multipart.FileHeader) (*os.File, error) {
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
