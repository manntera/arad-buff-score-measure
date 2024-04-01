package DetectedTextRepo

import (
	"errors"
	"image"
	"strings"

	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

type DetectedTextRepo struct {
	detectedText []*DetectedText
}

var _ IDetectedTextRepo = &DetectedTextRepo{}

func NewDetectedTextRepoFromVisionAnnotations(annotations []*visionpb.EntityAnnotation) (*DetectedTextRepo, error) {
	if len(annotations) == 0 {
		return nil, errors.New("annotations are empty")
	}
	result := &DetectedTextRepo{}
	detectedTexts := make([]*DetectedText, 0)
	lines := strings.Split(annotations[0].Description, "\n")

	annotationIndex := 0
	for _, line := range lines {
		if annotationIndex >= len(annotations) {
			break
		}
		detectedText := &DetectedText{
			text: line,
			rect: image.Rectangle{
				Min: image.Point{
					X: int(annotations[annotationIndex].BoundingPoly.Vertices[0].X),
					Y: int(annotations[annotationIndex].BoundingPoly.Vertices[0].Y),
				},
				Max: image.Point{
					X: int(annotations[annotationIndex].BoundingPoly.Vertices[2].X),
					Y: int(annotations[annotationIndex].BoundingPoly.Vertices[2].Y),
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
		}
		detectedTexts = append(detectedTexts, detectedText)
	}
	result.detectedText = detectedTexts
	return result, nil
}

func (repo *DetectedTextRepo) FindLineTextFromKeyword(keyword string, searchRect image.Rectangle) ([]*DetectedText, error) {
	result := make([]*DetectedText, 0)
	for _, detectedText := range repo.detectedText {
		if detectedText.rect.In(searchRect) && strings.Contains(detectedText.text, keyword) {
			result = append(result, detectedText)
		}
	}
	return result, nil
}
