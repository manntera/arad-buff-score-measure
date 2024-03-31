package SamplerImageRepo

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
)

type SamplerImageRepo struct {
	samplerImage *SamplerImage
}

var _ IImageSamplerRepo = &SamplerImageRepo{}

func NewSamplerImageRepoFromFileHeader(fileHeader *multipart.FileHeader) (*SamplerImageRepo, error) {
	result := &SamplerImageRepo{}

	samplerImage := SamplerImage{
		srcFile:  nil,
		srcImage: nil,
	}

	file, err := createFileFromFileHeader(fileHeader)
	if err != nil {
		return nil, err
	}

	image, err := getImageFromFile(PNG, file)
	if err != nil {
		file.Close()
		return nil, err
	}

	samplerImage.srcImage = &image
	samplerImage.srcFile = file
	result.samplerImage = &samplerImage

	return result, nil
}

func NewSamplerImageRepoFromFile(file *os.File) (*SamplerImageRepo, error) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	result := &SamplerImageRepo{}

	samplerImage := SamplerImage{
		srcFile:  nil,
		srcImage: nil,
	}

	image, err := getImageFromFile(PNG, file)
	if err != nil {
		return nil, err
	}

	samplerImage.srcImage = &image
	samplerImage.srcFile = file
	result.samplerImage = &samplerImage

	return result, nil
}

func (s *SamplerImageRepo) Close() error {
	if s.samplerImage.srcFile != nil {
		if err := s.samplerImage.srcFile.Close(); err != nil {
			return err
		}
		s.samplerImage.srcFile = nil
	}
	s.samplerImage.srcImage = nil
	return nil
}

func (s *SamplerImageRepo) GetFile() *os.File {
	if s.samplerImage.srcFile != nil {
		s.samplerImage.srcFile.Seek(0, io.SeekStart)
	}
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
	src, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	tempFile, err := os.CreateTemp("", "image")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(tempFile, src)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, err
	}

	_, err = tempFile.Seek(0, 0)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, err
	}

	return tempFile, nil
}

func getImageFromFile(imageType ImageType, imageFile *os.File) (image.Image, error) {
	switch imageType {
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
