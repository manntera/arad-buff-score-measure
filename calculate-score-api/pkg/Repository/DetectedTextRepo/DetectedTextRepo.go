package DetectedTextRepo

import (
	"context"
	"errors"
	"image"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"manntera.com/calculate-score-api/pkg/NormalizeRect"
	"manntera.com/calculate-score-api/pkg/Repository/SamplerImageRepo"
)

type DetectedTextRepo struct {
	detectedText []*DetectedText
}

var _ IDetectedTextRepo = &DetectedTextRepo{}

func NewDetectedTextRepoFromSamplerImageRepo(ctx context.Context, samplerImageRepo *SamplerImageRepo.SamplerImageRepo) (*DetectedTextRepo, error) {
	visionImage, imageErr := vision.NewImageFromReader(samplerImageRepo.GetFile())
	if imageErr != nil {
		return nil, imageErr
	}
	client, visionErr := vision.NewImageAnnotatorClient(ctx)
	if visionErr != nil {
		return nil, visionErr
	}
	defer client.Close()

	annotations, annotateErr := client.DetectTexts(ctx, visionImage, nil, 500)
	if annotateErr != nil {
		return nil, annotateErr
	}

	if len(annotations) == 0 {
		return nil, errors.New("annotations are empty")
	}
	result := &DetectedTextRepo{}
	detectedTexts := make([]*DetectedText, 0)
	lines := strings.Split(annotations[0].Description, "\n")

	srcImageRect := samplerImageRepo.GetImageSize()
	annotationIndex := 0
	for _, line := range lines {
		if annotationIndex >= len(annotations) {
			break
		}
		minPoint := image.Point{
			X: int(annotations[annotationIndex].BoundingPoly.Vertices[0].X),
			Y: int(annotations[annotationIndex].BoundingPoly.Vertices[0].Y),
		}
		maxPoint := image.Point{
			X: int(annotations[annotationIndex].BoundingPoly.Vertices[2].X),
			Y: int(annotations[annotationIndex].BoundingPoly.Vertices[2].Y),
		}
		detectedText := &DetectedText{
			text: line,
			rect: image.Rectangle{
				Min: minPoint,
				Max: maxPoint,
			},
			normalizeRect: NormalizeRect.NormalizeRect{
				Min: NormalizeRect.NormalizedPoint{
					X: float64(minPoint.X) / float64(srcImageRect.Width),
					Y: float64(minPoint.Y) / float64(srcImageRect.Height),
				},
				Max: NormalizeRect.NormalizedPoint{
					X: float64(maxPoint.X) / float64(srcImageRect.Width),
					Y: float64(maxPoint.Y) / float64(srcImageRect.Height),
				},
			},
		}
		for ; annotationIndex < len(annotations); annotationIndex++ {
			if detectedText.text == line {
				break
			}
			if annotationIndex >= len(annotations) {
				break
			}
			annotation := annotations[annotationIndex]
			detectedText.text += annotation.Description
			detectedText.rect.Max.X = int(annotation.BoundingPoly.Vertices[2].X)
			detectedText.rect.Max.Y = int(annotation.BoundingPoly.Vertices[2].Y)
			detectedText.normalizeRect.Max.X = float64(detectedText.rect.Max.X) / float64(srcImageRect.Width)
			detectedText.normalizeRect.Max.Y = float64(detectedText.rect.Max.Y) / float64(srcImageRect.Height)
		}
		detectedTexts = append(detectedTexts, detectedText)
	}
	result.detectedText = detectedTexts
	return result, nil
}

func (repo *DetectedTextRepo) FindLineTextFromKeyword(keyword string, searchNormalizeRect NormalizeRect.NormalizeRect) ([]*DetectedText, error) {
	result := make([]*DetectedText, 0)
	for _, detectedText := range repo.detectedText {
		if NormalizeRect.IsCollisionRect(detectedText.normalizeRect, searchNormalizeRect) {
			if strings.Contains(detectedText.text, keyword) {
				result = append(result, detectedText)
			}
		}
	}
	return result, nil
}
